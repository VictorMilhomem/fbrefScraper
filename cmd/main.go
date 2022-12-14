package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/VictorMilhomem/fbreScraper/cmd/colors"
	"github.com/gocolly/colly"
)

// TODO: url for custom teams "https://fbref.com/en/squads/ID/YEAR/all_comps/TEAMNAME-Stats-All-Competitions"
// TODO: maybe url for get comp tables

type ShootingStats struct {
	gls       string `json:"gols"`
	sht       string `json:"total_shots"`
	sot       string `json:"shots_on_target"`
	sot_per   string `json:"shots_on_target_per"`
	sht_ft    string `json:"shots_fulltime"`
	sot_ft    string `json:"shots_on_target_fulltime"`
	gls_shot  string `json:"gols_per_shot"`
	dist      string `json:"distance"`
	fk        string `json:"free_kick"`
	pk        string `json:"penalty_kick"`
	pkAtt     string `json:"penalty_kick_attempted"`
	xg        string `json:"xG"`
	npxg      string `json:"npxG"`
	npxg_shot string `json:"npxG_shot"`
	gls_xg    string `json:"G-xG"`
	np        string `json:"np:G-xG"`
}

type PassingStats struct {
}

type PassTypesStats struct {
}

type GoalShotCreationStats struct {
}

type DefenseStats struct {
}

type PossessionStats struct {
}

type PlayingTimeStats struct {
}

type MiscellaneousStats struct {
}

type PlayerSummaryStats struct {
}

type GoalKeepingStats struct {
	goalAgainst                 string `json:"goal_against"`
	penaltyKickAllowed          string `json:"pk_allowed"`
	freeKickAllowed             string `json:"fk_allowed"`
	cornerGoalsAgainst          string `json:"ck_allowed"`
	ownGoals                    string `json:"own_goals"`
	postShotxG                  string `json:"ps_xG"`
	postShotxGSot               string `json:"ps_xG_SoT"`
	postShotxGMinusGlsAllowed   string `json:"ps_xG_minus_gls_allowed"`
	postShotxGMinusGlsAllowedFt string `json:"ps_xG_minus_gls_allowed_ft"`
	// launched more than 40yards
	launchedPassesCompletedLonger string `json:"passes_cmp_longer"`
	launchedPassesAttemptedLonger string `json:"passes_att_longer"`
	launchedPassesCompletedPer    string `json:"passes_cmp_per"`
	// passes not including goal kick
	passesAttempted string `json:"passes_att"`
	throwsAttempted string `json:"thr_att"`
	launchPer       string `json:"launch_per"`
	averageLen      string `json:"avg_len"`
	//goal kicks
	gkAttempted  string `json:"gk_att"`
	gkLaunchPer  string `json:"gk_launch_per"`
	gkAverageLen string `json:"gk_avg_len"`
	//Crosses
	opponentAttemptedCrosses string `json:"crosses_opp"`
	crossesStopedByKeeper    string `json:"crosses_stp_by_keeper"`
	crossesStopedByKeeperPer string `json:"crosses_stp_by_keeper_per"`
	// sweeper
	nDefensiveActionsByKeeper   string `json:"n_defense_actions_by_keeper"`
	nDefensiveActionsByKeeperFt string `json:"n_defense_actions_by_keeper_ft"`
	averageDistance             string `json:"average_distance"`
}

//TODO: maybe create a interface for a player, to solve the problem where i need to get the stats from diff tables
// TODO: create a struct to handle all players statistics
type Player struct {
	name             string                `json:"name"`
	nation           string                `json:"nation"`
	pos              string                `json:"position"`
	age              string                `json:"age"`
	min              string                `json:"minutes_played"`
	shooting         ShootingStats         `json: "shooting_stats"`
	passing          PassingStats          `json:"passing_stats"`
	passType         PassTypesStats        `json:"passing_types"`
	goalShotCreation GoalShotCreationStats `json:"goal_shot_creation"`
	defenseActions   DefenseStats          `json:"defense_actions"`
	possession       PossessionStats       `json:"possession"`
	playingTime      PlayingTimeStats      `json:"playing_time"`
	miscellaneous    MiscellaneousStats    `json:"miscellaneous"`
	playerSummary    PlayerSummaryStats    `json:"player_summary"`
	goalkeeping      GoalKeepingStats      `json:"goalkeeping"`
}

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
	var players []Player
	c.OnHTML("table#stats_shooting_combined > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, el *colly.HTMLElement) {
			player := Player{
				name:   el.ChildText("th"),
				nation: el.ChildText("td:nth-child(1)"),
				pos:    el.ChildText("td:nth-child(2)"),
				age:    el.ChildText("td:nth-child(3)"),
				min:    el.ChildText("td:nth-child(4)"),
				shooting: ShootingStats{
					gls:       el.ChildText("td:nth-child(5)"),
					sht:       el.ChildText("td:nth-child(6)"),
					sot:       el.ChildText("td:nth-child(7)"),
					sot_per:   el.ChildText("td:nth-child(8)"),
					sht_ft:    el.ChildText("td:nth-child(9)"),
					sot_ft:    el.ChildText("td:nth-child(10)"),
					gls_shot:  el.ChildText("td:nth-child(11)"),
					dist:      el.ChildText("td:nth-child(12)"),
					fk:        el.ChildText("td:nth-child(13)"),
					pk:        el.ChildText("td:nth-child(14)"),
					pkAtt:     el.ChildText("td:nth-child(15)"),
					xg:        el.ChildText("td:nth-child(16)"),
					npxg:      el.ChildText("td:nth-child(17)"),
					npxg_shot: el.ChildText("td:nth-child(18)"),
					gls_xg:    el.ChildText("td:nth-child(19)"),
					np:        el.ChildText("td:nth-child(20)"),
				},
			}
			// p, err := json.Marshal(player)
			// if err != nil {
			//     log.Fatal(err)
			// }
			players = append(players, player)
			log.Println(players)
		})
	})

	c.Visit("https://fbref.com/en/squads/84d9701c/2022/all_comps/Fluminense-Stats-All-Competitions")

}
