package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/VictorMilhomem/fbreScraper/cmd/colors"
	"github.com/VictorMilhomem/fbreScraper/cmd/types"
	"github.com/gocolly/colly"
)

// TODO: url for custom teams "https://fbref.com/en/squads/ID/YEAR/all_comps/TEAMNAME-Stats-All-Competitions"
// TODO: maybe url for get comp tables
var playerInfo types.PlayerBasic

var (
	shootingStats types.ShootingStats
	player        *types.Player
)

// TODO: Accept all teams stastistics and players by id and name ?

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("fbref.com"),
	)

	c.WithTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   90 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println(colors.Cyan, "Scraping:", r.URL, colors.Reset)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(colors.Green, "Status:", r.StatusCode, colors.Reset)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println(colors.Red, "Request URL:", colors.Reset, r.Request.URL, "failed with response:", r, colors.Red, "\nError:", err, colors.Reset)
	})

	// TODO: getting the struct field infos from each respective table

	// Shooting statistics
	var players []types.Player
	c.OnHTML("table#stats_shooting_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			playerInfo = types.PlayerBasic{
				Name:   el.ChildText("th"),
				Nation: el.ChildText("td:nth-child(2)"),
				Pos:    el.ChildText("td:nth-child(3)"),
				Age:    el.ChildText("td:nth-child(4)"),
				Min:    el.ChildText("td:nth-child(5)"),
			}

			shootingStats = types.ShootingStats{
				Gls:         el.ChildText("td:nth-child(6)"),
				Sht:         el.ChildText("td:nth-child(7)"),
				Sot:         el.ChildText("td:nth-child(8)"),
				Sot_per:     el.ChildText("td:nth-child(9)"),
				Sht_ft:      el.ChildText("td:nth-child(10)"),
				Sot_ft:      el.ChildText("td:nth-child(11)"),
				Gls_shot:    el.ChildText("td:nth-child(12)"),
				Gls_per_sot: el.ChildText("td:nth-child(13)"),
				Dist:        el.ChildText("td:nth-child(14)"),
				Fk:          el.ChildText("td:nth-child(15)"),
				Pk:          el.ChildText("td:nth-child(16)"),
				PkAtt:       el.ChildText("td:nth-child(17)"),
				Xg:          el.ChildText("td:nth-child(18)"),
				Npxg:        el.ChildText("td:nth-child(19)"),
				Npxg_shot:   el.ChildText("td:nth-child(20)"),
				Gls_xg:      el.ChildText("td:nth-child(21)"),
				Np:          el.ChildText("td:nth-child(22)"),
			}

			//player = types.Player{
			//	PlayerBasicInfo: playerInfo,
			//	Shooting:        shootingStats,
			//}
			player = types.NewPlayer(playerInfo)
			player.AppendShooting(shootingStats)
			players = append(players, *player)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(players[1])
	})

	c.Visit("https://fbref.com/en/squads/84d9701c/2022/all_comps/Fluminense-Stats-All-Competitions")
}
