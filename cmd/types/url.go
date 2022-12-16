package types

type Url struct {
	ID    string
	Team  string
	Value string
}

func NewUrlTeamStats(id string, team string, year string) *Url {
	return &Url{
		ID:    id,
		Team:  team,
		Value: "https://fbref.com/en/squads/" + id + "/" + year + "/all_comps/" + team + "-Stats-All-Competitions",
	}
}
