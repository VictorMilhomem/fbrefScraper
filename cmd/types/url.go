package types

import "strconv"

type Url struct {
	ID    string
	Team  string
	Value string
}

func NewUrlTeamStats(id string, team string, year int) *Url {
	return &Url{
		ID:    id,
		Team:  team,
		Value: "https://fbref.com/en/squads/" + id + "/" + strconv.Itoa(year) + "/all_comps/" + team + "-Stats-All-Competitions",
	}
}
