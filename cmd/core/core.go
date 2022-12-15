package core

import (
	"encoding/json"
	"log"
	"os"

	"github.com/VictorMilhomem/fbreScraper/cmd/types"
	"github.com/gocolly/colly"
)

var players []types.Player

// TODO: create functions to get all players statistics

func getPlayersShootingStats(c *colly.Collector) error {
	var (
		playerInfo    types.PlayerBasic
		shootingStats types.ShootingStats
		player        *types.Player
	)

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

			player = types.NewPlayer(playerInfo)
			player.AppendShooting(shootingStats)
			players = append(players, *player)
		})
	})

	// Shooting statistics
	return nil
}

func ScrapePlayers(url string, c *colly.Collector) {
	// TODO: getting the struct field infos from each respective table

	err := getPlayersShootingStats(c)
	if err != nil {
		log.Fatal("Error getting players shooting statistics: ", err)
	}

	c.OnScraped(func(r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(players)
	})

	c.Visit(url)
}
