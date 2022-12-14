package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
	"unicode"

	api "github.com/VictorMilhomem/fbreScraper/cmd/api/server"
	"github.com/VictorMilhomem/fbreScraper/cmd/colors"
	"github.com/VictorMilhomem/fbreScraper/cmd/core"
	"github.com/VictorMilhomem/fbreScraper/cmd/types"
	"github.com/gocolly/colly"
)

// TODO: url for custom teams "https://fbref.com/en/squads/ID/YEAR/all_comps/TEAMNAME-Stats-All-Competitions"
// TODO: maybe url for get comp tables
// TODO: Accept all teams stastistics and players by id and name ?

func collyConfig(c *colly.Collector) {
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
}

func UpperCaseFirstChar(str string) string {
	str = strings.ToLower(str)
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}

	return ""
}

func main() {
	fmt.Println(colors.Cyan, "Welcome to fbref.com data scraper")

	id := flag.String("id", "", "set a team id at fbref.com")
	team := flag.String("t", "", "set an avaible team at fbref.com")
	year := flag.Int("y", 2022, "set the year from avaibles years at fbref.com")
	serve := flag.Bool("s", false, "start server")

	flag.Parse()

	if *serve {
		api.StartServer()
	}

	var url types.Url
	if *id != "" && *team != "" {
		url = *types.NewUrlTeamStats(*id, UpperCaseFirstChar(*team), *year)
		// url := "https://fbref.com/en/squads/84d9701c/2022/all_comps/Fluminense-Stats-All-Competitions"
		c := colly.NewCollector(
			colly.AllowedDomains("fbref.com"),
		)
		collyConfig(c)
		core.ScrapePlayers(url, c)
	} else {
		log.Fatal(colors.Red, "Unknown flag:", colors.Reset)
	}
}
