package core

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/VictorMilhomem/fbreScraper/cmd/colors"
	"github.com/VictorMilhomem/fbreScraper/cmd/types"
	"github.com/gocolly/colly"
)

var players []types.Player

// TODO: create functions to get all players statistics

func getPassingStats(c *colly.Collector) {
	var passing *types.PassingStats

	c.OnHTML("table#stats_passing_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			passing = types.NewPassingStats(
				el.ChildText("td:nth-child(6)"),
				el.ChildText("td:nth-child(7)"),
				el.ChildText("td:nth-child(8)"),
				el.ChildText("td:nth-child(9)"),
				el.ChildText("td:nth-child(10)"),
				el.ChildText("td:nth-child(11)"),
				el.ChildText("td:nth-child(12)"),
				el.ChildText("td:nth-child(13)"),
				el.ChildText("td:nth-child(14)"),
				el.ChildText("td:nth-child(15)"),
				el.ChildText("td:nth-child(16)"),
				el.ChildText("td:nth-child(17)"),
				el.ChildText("td:nth-child(18)"),
				el.ChildText("td:nth-child(19)"),
				el.ChildText("td:nth-child(20)"),
				el.ChildText("td:nth-child(21)"),
				el.ChildText("td:nth-child(22)"),
				el.ChildText("td:nth-child(23)"),
				el.ChildText("td:nth-child(24)"),
				el.ChildText("td:nth-child(25)"),
				el.ChildText("td:nth-child(26)"),
				el.ChildText("td:nth-child(27)"),
				el.ChildText("td:nth-child(28)"),
			)

			players[i].AppendPassingStats(*passing)
		})
	})
}

func getGkAdvancedStats(c *colly.Collector) {
	var gkStats *types.GoalKeepingStats

	c.OnHTML("table#stats_keeper_adv_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			gkStats = types.NewGoalKeepingStats(
				el.ChildText("td:nth-child(6)"),
				el.ChildText("td:nth-child(7)"),
				el.ChildText("td:nth-child(8)"),
				el.ChildText("td:nth-child(9)"),
				el.ChildText("td:nth-child(10)"),
				el.ChildText("td:nth-child(11)"),
				el.ChildText("td:nth-child(12)"),
				el.ChildText("td:nth-child(13)"),
				el.ChildText("td:nth-child(14)"),
				el.ChildText("td:nth-child(15)"),
				el.ChildText("td:nth-child(16)"),
				el.ChildText("td:nth-child(17)"),
				el.ChildText("td:nth-child(18)"),
				el.ChildText("td:nth-child(19)"),
				el.ChildText("td:nth-child(20)"),
				el.ChildText("td:nth-child(21)"),
				el.ChildText("td:nth-child(22)"),
				el.ChildText("td:nth-child(23)"),
				el.ChildText("td:nth-child(24)"),
				el.ChildText("td:nth-child(25)"),
				el.ChildText("td:nth-child(26)"),
				el.ChildText("td:nth-child(27)"),
				el.ChildText("td:nth-child(28)"),
				el.ChildText("td:nth-child(29)"),
				el.ChildText("td:nth-child(30)"),
			)

			players[i].AppendGoalKeeping(*gkStats)
		})
	})
}

func getPlayersShootingStats(c *colly.Collector) {
	var (
		playerInfo    *types.PlayerBasic
		shootingStats *types.ShootingStats
		player        *types.Player
	)

	c.OnHTML("table#stats_shooting_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			playerInfo = types.NewPlayerBasic(
				el.ChildText("th"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
				el.ChildText("td:nth-child(4)"),
				el.ChildText("td:nth-child(5)"),
			)

			shootingStats = types.NewShootingStats(
				el.ChildText("td:nth-child(6)"),
				el.ChildText("td:nth-child(7)"),
				el.ChildText("td:nth-child(8)"),
				el.ChildText("td:nth-child(9)"),
				el.ChildText("td:nth-child(10)"),
				el.ChildText("td:nth-child(11)"),
				el.ChildText("td:nth-child(12)"),
				el.ChildText("td:nth-child(13)"),
				el.ChildText("td:nth-child(14)"),
				el.ChildText("td:nth-child(15)"),
				el.ChildText("td:nth-child(16)"),
				el.ChildText("td:nth-child(17)"),
				el.ChildText("td:nth-child(18)"),
				el.ChildText("td:nth-child(19)"),
				el.ChildText("td:nth-child(20)"),
				el.ChildText("td:nth-child(21)"),
				el.ChildText("td:nth-child(22)"),
			)

			player = types.NewPlayer(*playerInfo)
			player.AppendShooting(*shootingStats)
			players = append(players, *player)
		})
	})

	// Shooting statistics
}

func writeFile(filename string, p types.Players) {
	file, _ := json.MarshalIndent(p, "", " ")
	_ = ioutil.WriteFile(filename, file, 0o644)
	log.Println(colors.Green, "File ", colors.Cyan, filename, colors.Green, " created", colors.Reset)
}

func ScrapePlayers(url string, c *colly.Collector) {
	// TODO: getting the struct field infos from each respective table

	getPlayersShootingStats(c)
	getPassingStats(c)
	getGkAdvancedStats(c)

	c.OnScraped(func(r *colly.Response) {
		// enc := json.NewEncoder(&os.File{})
		// enc.Encode(players)
		// scrape team name
		p := *types.NewPlayers("fluminense", players)
		filename := "players_stats.json"
		writeFile(filename, p)
	})

	c.Visit(url)
}
