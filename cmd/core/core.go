package core

import (
	"encoding/json"
	"io/ioutil"

	"github.com/VictorMilhomem/fbreScraper/cmd/types"
	"github.com/gocolly/colly"
)

var players []types.Player

// TODO: create functions to get all players statistics

func getPassingStats(c *colly.Collector) {
	var passing types.PassingStats

	c.OnHTML("table#stats_passing_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			passing = types.PassingStats{
				PassesCompleted:                  el.ChildText("td:nth-child(6)"),
				PassesAttempted:                  el.ChildText("td:nth-child(7)"),
				CompletionPercent:                el.ChildText("td:nth-child(8)"),
				TotalDist:                        el.ChildText("td:nth-child(9)"),
				PrgDist:                          el.ChildText("td:nth-child(10)"),
				ShortPassesCompleted:             el.ChildText("td:nth-child(11)"),
				ShortPassesAttempted:             el.ChildText("td:nth-child(12)"),
				ShortPassesCmpPercent:            el.ChildText("td:nth-child(13)"),
				MediumPassesCompleted:            el.ChildText("td:nth-child(14)"),
				MediumPassesAttempted:            el.ChildText("td:nth-child(15)"),
				MediumPassesCmpPercent:           el.ChildText("td:nth-child(16)"),
				LongPassesCompleted:              el.ChildText("td:nth-child(17)"),
				LongPassesAttempted:              el.ChildText("td:nth-child(18)"),
				LongPassesCmpPercent:             el.ChildText("td:nth-child(19)"),
				Assists:                          el.ChildText("td:nth-child(20)"),
				XAG:                              el.ChildText("td:nth-child(21)"),
				XA:                               el.ChildText("td:nth-child(22)"),
				AssistsMinusXAG:                  el.ChildText("td:nth-child(23)"),
				AssistedShot:                     el.ChildText("td:nth-child(24)"),
				CompletedPassesThatEnterOneThird: el.ChildText("td:nth-child(25)"),
				CompletedPassesIntoTheBox:        el.ChildText("td:nth-child(26)"),
				CompletedCrossesIntoTheBox:       el.ChildText("td:nth-child(27)"),
				PrgPasses:                        el.ChildText("td:nth-child(28)"),
			}

			players[i].Passing = passing
		})
	})
}

func getGkAdvancedStats(c *colly.Collector) {
	var gkStats types.GoalKeepingStats

	c.OnHTML("table#stats_keeper_adv_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			gkStats = types.GoalKeepingStats{
				GoalAgainst:                 el.ChildText("td:nth-child(6)"),
				PenaltyKickAllowed:          el.ChildText("td:nth-child(7)"),
				FreeKickAllowed:             el.ChildText("td:nth-child(8)"),
				CornerGoalsAgainst:          el.ChildText("td:nth-child(9)"),
				OwnGoals:                    el.ChildText("td:nth-child(10)"),
				PostShotxG:                  el.ChildText("td:nth-child(11)"),
				PostShotxGSot:               el.ChildText("td:nth-child(12)"),
				PostShotxGMinusGlsAllowed:   el.ChildText("td:nth-child(13)"),
				PostShotxGMinusGlsAllowedFt: el.ChildText("td:nth-child(14)"),

				// launched more than 40yards
				LaunchedPassesCompletedLonger: el.ChildText("td:nth-child(15)"),
				LaunchedPassesAttemptedLonger: el.ChildText("td:nth-child(16)"),
				LaunchedPassesCompletedPer:    el.ChildText("td:nth-child(17)"),

				// passes not including goal kick
				PassesAttempted: el.ChildText("td:nth-child(18)"),
				ThrowsAttempted: el.ChildText("td:nth-child(19)"),
				LaunchPer:       el.ChildText("td:nth-child(20)"),
				AverageLen:      el.ChildText("td:nth-child(21)"),

				// goal kicks
				GkAttempted:  el.ChildText("td:nth-child(22)"),
				GkLaunchPer:  el.ChildText("td:nth-child(23)"),
				GkAverageLen: el.ChildText("td:nth-child(24)"),

				// Crosses
				OpponentAttemptedCrosses: el.ChildText("td:nth-child(25)"),
				CrossesStopedByKeeper:    el.ChildText("td:nth-child(26)"),
				CrossesStopedByKeeperPer: el.ChildText("td:nth-child(27)"),

				// sweeper
				NDefensiveActionsByKeeper:   el.ChildText("td:nth-child(28)"),
				NDefensiveActionsByKeeperFt: el.ChildText("td:nth-child(29)"),
				AverageDistance:             el.ChildText("td:nth-child(30)"),
			}

			players[i].Goalkeeping = gkStats
		})
	})
}

func getPlayersShootingStats(c *colly.Collector) {
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
}

func ScrapePlayers(url string, c *colly.Collector) {
	// TODO: getting the struct field infos from each respective table

	getPlayersShootingStats(c)
	getPassingStats(c)
	getGkAdvancedStats(c)

	c.OnScraped(func(r *colly.Response) {
		// enc := json.NewEncoder(&os.File{})
		// enc.Encode(players)
		file, _ := json.MarshalIndent(players, "", " ")
		_ = ioutil.WriteFile("players_stats.json", file, 0o644)
	})

	c.Visit(url)
}
