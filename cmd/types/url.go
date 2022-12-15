package types

type Url struct {
	Value string
}

func NewUrlTeamStats(id string, team string, year string) *Url {
	return &Url{
		Value: "https://fbref.com/en/squads/" + id + "/" + year + "/all_comps/" + team + "-Stats-All-Competitions",
	}
}
