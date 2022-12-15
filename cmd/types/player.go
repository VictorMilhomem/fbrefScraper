package types

type ShootingStats struct {
	Gls         string `json:"gols"`
	Sht         string `json:"total_shots"`
	Sot         string `json:"shots_on_target"`
	Sot_per     string `json:"shots_on_target_per"`
	Sht_ft      string `json:"shots_fulltime"`
	Sot_ft      string `json:"shots_on_target_fulltime"`
	Gls_shot    string `json:"gols_per_shot"`
	Gls_per_sot string `json:"gols_per_SoT"`
	Dist        string `json:"distance"`
	Fk          string `json:"free_kick"`
	Pk          string `json:"penalty_kick"`
	PkAtt       string `json:"penalty_kick_attempted"`
	Xg          string `json:"xG"`
	Npxg        string `json:"npxG"`
	Npxg_shot   string `json:"npxG_shot"`
	Gls_xg      string `json:"G-xG"`
	Np          string `json:"np:G-xG"`
}

type PassingStats struct{}

type PassTypesStats struct{}

type GoalShotCreationStats struct{}

type DefenseStats struct{}

type PossessionStats struct{}

type PlayingTimeStats struct{}

type MiscellaneousStats struct{}

type PlayerSummaryStats struct{}

type GoalKeepingStats struct {
	GoalAgainst                 string `json:"goal_against"`
	PenaltyKickAllowed          string `json:"pk_allowed"`
	FreeKickAllowed             string `json:"fk_allowed"`
	CornerGoalsAgainst          string `json:"ck_allowed"`
	OwnGoals                    string `json:"own_goals"`
	PostShotxG                  string `json:"ps_xG"`
	PostShotxGSot               string `json:"ps_xG_SoT"`
	PostShotxGMinusGlsAllowed   string `json:"ps_xG_minus_gls_allowed"`
	PostShotxGMinusGlsAllowedFt string `json:"ps_xG_minus_gls_allowed_ft"`

	// launched more than 40yards
	LaunchedPassesCompletedLonger string `json:"passes_cmp_longer"`
	LaunchedPassesAttemptedLonger string `json:"passes_att_longer"`
	LaunchedPassesCompletedPer    string `json:"passes_cmp_per"`

	// passes not including goal kick
	PassesAttempted string `json:"passes_att"`
	ThrowsAttempted string `json:"thr_att"`
	LaunchPer       string `json:"launch_per"`
	AverageLen      string `json:"avg_len"`

	// goal kicks
	GkAttempted  string `json:"gk_att"`
	GkLaunchPer  string `json:"gk_launch_per"`
	GkAverageLen string `json:"gk_avg_len"`

	// Crosses
	OpponentAttemptedCrosses string `json:"crosses_opp"`
	CrossesStopedByKeeper    string `json:"crosses_stp_by_keeper"`
	CrossesStopedByKeeperPer string `json:"crosses_stp_by_keeper_per"`

	// sweeper
	NDefensiveActionsByKeeper   string `json:"n_defense_actions_by_keeper"`
	NDefensiveActionsByKeeperFt string `json:"n_defense_actions_by_keeper_ft"`
	AverageDistance             string `json:"average_distance"`
}

// TODO: maybe create a interface for a player, to solve the problem where i need to get the stats from diff tables
// TODO: create a struct to handle all players statistics
type PlayerBasic struct {
	Name   string
	Nation string
	Pos    string
	Age    string
	Min    string
}

type Player struct {
	Name   string `json:"name"`
	Nation string `json:"nation"`
	Pos    string `json:"position"`
	Age    string `json:"age"`
	Min    string `json:"matches_played"`

	Shooting         ShootingStats         `json: "shooting_stats"`
	Passing          PassingStats          `json:"passing_stats"`
	PassType         PassTypesStats        `json:"passing_types"`
	GoalShotCreation GoalShotCreationStats `json:"goal_shot_creation"`
	DefenseActions   DefenseStats          `json:"defense_actions"`
	Possession       PossessionStats       `json:"possession"`
	PlayingTime      PlayingTimeStats      `json:"playing_time"`
	Miscellaneous    MiscellaneousStats    `json:"miscellaneous"`
	PlayerSummary    PlayerSummaryStats    `json:"player_summary"`
	Goalkeeping      GoalKeepingStats      `json:"goalkeeping"`
}

func NewPlayer(info PlayerBasic) *Player {
	return &Player{
		Name:   info.Name,
		Nation: info.Nation,
		Pos:    info.Pos,
		Age:    info.Age,
		Min:    info.Min,
	}
}

func (p *Player) AppendShooting(st ShootingStats) {
	p.Shooting = st
}

func (p *Player) AppendGoalKeeping(gk GoalKeepingStats) {
	p.Goalkeeping = gk
}
