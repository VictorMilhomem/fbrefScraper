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

// TODO: getPlayerSummary current not working
func getPlayerSummary(c *colly.Collector) {
	var summary *types.PlayerSummaryStats
	c.OnHTML("table#stats_player_summary > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			summary = types.NewPlayerSummary(
				el.ChildText("td:nth-child(13)"),
				el.ChildText("td:nth-child(14)"),
				el.ChildText("td:nth-child(15)"),
				el.ChildText("td:nth-child(16)"),
			)
			if i < len(players) {
				players[i].AppendPlayerSummary(*summary)
			}
		})
	})
}

func getMiscellaneousStats(c *colly.Collector) {
	var mis *types.MiscellaneousStats
	c.OnHTML("table#stats_misc_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			mis = types.NewMiscellaneousStats(
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
			)
			players[i].AppendMisscellaneous(*mis)
		})
	})
}

func getPlayingTimeStats(c *colly.Collector) {
	var playingTime *types.PlayingTimeStats
	c.OnHTML("table#stats_playing_time_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			playingTime = types.NewPlayingTimeStats(
				el.ChildText("td:nth-child(5)"),
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
			)
			if i < len(players) {
				players[i].AppendPlayingTime(*playingTime)
			}
		})
	})
}

func getPossessionStats(c *colly.Collector) {
	var possession *types.PossessionStats
	c.OnHTML("table#stats_possession_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			possession = types.NewPossessionStats(
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
			)
			players[i].AppendPossession(*possession)
		})
	})
}

func getDefensiveActionsStats(c *colly.Collector) {
	var defense *types.DefenseStats
	c.OnHTML("table#stats_defense_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			defense = types.NewDefenseStats(
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
			)
			players[i].AppendDefensiveActions(*defense)
		})
	})
}

func getGoalShotStats(c *colly.Collector) {
	var gs *types.GoalShotCreationStats
	c.OnHTML("table#stats_gca_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			gs = types.NewGoalShotCreation(
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
			)
			players[i].AppendGoalCreation(*gs)
		})
	})
}

func getPassingTypeStats(c *colly.Collector) {
	var passtypes *types.PassTypesStats

	c.OnHTML("table#stats_passing_types_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			passtypes = types.NewPassTypes(
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
			)
			players[i].AppendPassTypes(*passtypes)
		})
	})
}

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
	getGkAdvancedStats(c)
	getPassingStats(c)
	getPassingTypeStats(c)
	getGoalShotStats(c)
	getDefensiveActionsStats(c)
	getPossessionStats(c)
	getPlayingTimeStats(c)
	getMiscellaneousStats(c)
	getPlayerSummary(c)

	c.OnScraped(func(r *colly.Response) {
		// scrape team name
		p := *types.NewPlayers("fluminense", players)
		filename := "players_stats.json"
		writeFile(filename, p)
	})

	c.Visit(url)
}
