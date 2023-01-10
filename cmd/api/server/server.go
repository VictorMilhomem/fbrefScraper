package api

import (
	"net/http"

	"github.com/VictorMilhomem/fbreScraper/cmd/api/handler"
)

func StartServer() {
	http.HandleFunc("/teams", handler.PlayersHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
