package main

import (
	"log"
	"net"
	"net/http"
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
	player        types.Player
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
				Nation: el.ChildText("td:nth-child(1)"),
				Pos:    el.ChildText("td:nth-child(2)"),
				Age:    el.ChildText("td:nth-child(3)"),
				Min:    el.ChildText("td:nth-child(4)"),
			}

			shootingStats = types.ShootingStats{
				Gls:       el.ChildText("td:nth-child(5)"),
				Sht:       el.ChildText("td:nth-child(6)"),
				Sot:       el.ChildText("td:nth-child(7)"),
				Sot_per:   el.ChildText("td:nth-child(8)"),
				Sht_ft:    el.ChildText("td:nth-child(9)"),
				Sot_ft:    el.ChildText("td:nth-child(10)"),
				Gls_shot:  el.ChildText("td:nth-child(11)"),
				Dist:      el.ChildText("td:nth-child(12)"),
				Fk:        el.ChildText("td:nth-child(13)"),
				Pk:        el.ChildText("td:nth-child(14)"),
				PkAtt:     el.ChildText("td:nth-child(15)"),
				Xg:        el.ChildText("td:nth-child(16)"),
				Npxg:      el.ChildText("td:nth-child(17)"),
				Npxg_shot: el.ChildText("td:nth-child(18)"),
				Gls_xg:    el.ChildText("td:nth-child(19)"),
				Np:        el.ChildText("td:nth-child(20)"),
			}
			// p, err := json.Marshal(player)
			// if err != nil {
			//     log.Fatal(err)
			// }
			player = types.Player{
				PlayerBasicInfo: playerInfo,
				Shooting:        shootingStats,
			}
			players = append(players, player)
			log.Println(players)
		})
	})

	c.Visit("https://fbref.com/en/squads/84d9701c/2022/all_comps/Fluminense-Stats-All-Competitions")
}
