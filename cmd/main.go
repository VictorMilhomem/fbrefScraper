package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/VictorMilhomem/fbreScraper/cmd/colors"
	"github.com/VictorMilhomem/fbreScraper/cmd/core"
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

func main() {
	url := "https://fbref.com/en/squads/84d9701c/2022/all_comps/Fluminense-Stats-All-Competitions"
	c := colly.NewCollector(
		colly.AllowedDomains("fbref.com"),
	)
	collyConfig(c)
	core.ScrapePlayers(url, c)
}
