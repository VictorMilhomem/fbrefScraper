package main

import (
	"net"
	"net/http"
	"time"
    "log"
	"github.com/gocolly/colly"
)


// TODO: create a struct to handle all players statistics
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
		log.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Status:", r.StatusCode)
	})
  

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

    c.OnHTML("table#stats_shooting_combined > tbody", func(h *colly.HTMLElement) {
        h.ForEach("tr", func(i int, el *colly.HTMLElement) {
            log.Println(el.ChildText("th"))
        }) 
	}) 

	c.Visit("https://fbref.com/en/squads/84d9701c/2022/all_comps/Fluminense-Stats-All-Competitions")


}
