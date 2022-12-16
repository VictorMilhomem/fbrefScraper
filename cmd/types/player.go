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

type PassingStats struct {
	PassesCompleted                  string `json:"passes_completed"`
	PassesAttempted                  string `json:"passes_att"`
	CompletionPercent                string `json:"passes_cmp_percent"`
	TotalDist                        string `json:"passes_total_dist"`
	PrgDist                          string `json:"progressive_dist"`
	ShortPassesCompleted             string `json:"short_passes_cmp"`
	ShortPassesAttempted             string `json:"short_passes_att"`
	ShortPassesCmpPercent            string `json:"short_passes_cmp_percent"`
	MediumPassesCompleted            string `json:"medium_passes_cmp"`
	MediumPassesAttempted            string `json:"medium_passes_att"`
	MediumPassesCmpPercent           string `json:"medium_passes_cmp_percent"`
	LongPassesCompleted              string `json:"Long_passes_cmp"`
	LongPassesAttempted              string `json:"Long_passes_att"`
	LongPassesCmpPercent             string `json:"Longm_passes_cmp_percent"`
	Assists                          string `json:"assists"`
	XAG                              string `json:"xAG"`
	XA                               string `json:"xA"`
	AssistsMinusXAG                  string `json:"A-xAG"`
	AssistedShot                     string `json:"assisted_shot"`
	CompletedPassesThatEnterOneThird string `json:"cmp_passes_1/3"`
	CompletedPassesIntoTheBox        string `json:"ppa"`
	CompletedCrossesIntoTheBox       string `json:"crs_pa"`
	PrgPasses                        string `json:"progressive_passes"`
}

type PassTypesStats struct {
	PassesAttempted                 string `json:"passes_att"`
	LiveBallPasses                  string `json:"live_ball_passes"`
	DeadBallPasses                  string `json:"dead_ball_passes"`
	PassesAttemptedFromFreeKicks    string `json:"passes_att_fk"`
	CompletedPassesIntoOpenSpace    string `json:"passes_completed_tb"`
	PassesThatTravelMoreThan40yards string `json:"passes_that_travel_40yards+"`
	Crosses                         string `json:"crosses"`
	ThrowInTaken                    string `json:"throw_in"`
	CornerKicks                     string `json:"corner_kicks"`
	InSwingingCornerKicks           string `json:"inswinging_ck"`
	OutSwingingCornerKicks          string `json:"outswinging_ck"`
	StraightCornerKicks             string `json:"straitght_ck"`
	OutComesPassesCompleted         string `json:"outcomes_passes_completed"`
	OffsidePasses                   string `json:"offiside_passes"`
	PassesBlockedByOpp              string `json:"passes_blocked_by_opp"`
}

type GoalShotCreationStats struct {
	ShotCreatingActions              string `json:"shot_creating_actions"`
	ShotCreatingActionsPer90         string `json:"shot_creating_actions_per90"`
	SCPassLiveBall                   string `json:"sc_pass_live_ball"`
	SCPassDeadBall                   string `json:"sc_pass_dead_ball"`
	SCSuccessDribleLeadShot          string `json:"sc_drible_lead_shot"`
	SCShotThatLeadAnotherShot        string `json:"sc_shot_that_lead_another_shot"`
	SCFoulsDrawnThatLeadToShot       string `json:"sc_fouls_drawn_that_lead_shot"`
	SCDefensiveActionsThatLeadToShot string `json:"sc_defensive_actions_that_lead_shot"`
	GoalCreationAction               string `json:"goal_creation_action"`
	GoalCreationActionPer90          string `json:"goal_creation_action_per90"`
	GCPassLiveBall                   string `json:"gc_pass_live_ball"`
	GCPassDeadBall                   string `json:"gc_pass_dead_ball"`
	GCSuccessDribleLeadShot          string `json:"gc_drible_lead_shot"`
	GCShotThatLeadAnotherShot        string `json:"gc_shot_that_lead_another_shot"`
	GCFoulsDrawnThatLeadToShot       string `json:"gc_fouls_drawn_that_lead_shot"`
	GCDefensiveActionsThatLeadToShot string `json:"gc_defensive_actions_that_lead_shot"`
}

type DefenseStats struct {
	Tackles                   string `json:"tackles"`
	TacklesWonPossession      string `json:"tackles_won_possesion"`
	TacklesDefensive3rd       string `json:"tackles_defensive_3rd"`
	TacklesMid3rd             string `json:"tackles_mid_3rd"`
	TacklesAttacking3rd       string `json:"tackles_attacking_3rd"`
	NumberOfDribblersTackled  string `json:"n_dribblers_tackled"`
	NumberOfTimesDribblerPast string `json:"n_dribblers_past"`
	PercentOfDribblersTackled string `json:"percent_of_dribblers_tackled"`
	Blocks                    string `json:"blocks"`
	BlockedShot               string `json:"blocked_shot"`
	BlockkedPass              string `json:"blocked_pass"`
	Interceptions             string `json:"interception"`
	TacklePlusInterception    string `json:"tackle_plus_interception"`
	Clearances                string `json:"clearances"`
	ErrorLeadShot             string `json:"error_lead_shot"`
}

type PossessionStats struct {
	Touches                 string `json:"touches"`
	TouchesDefBoxArea       string `json:"touches_def_box"`
	TouchesDef3rd           string `json:"touches_def_3rd"`
	TouchesMid3rd           string `json:"touches_mid_3rd"`
	TouchesAttacking3rd     string `json:"touches_attacking_3rd"`
	TouchesAttackingBoxArea string `json:"touches_attacking_box"`
	LiveBallTouches         string `json:"touches_live_ball"`
	DribblesSucc            string `json:"success_dribbles"`
	DribblesAttempted       string `json:"dribbles_attempted"`
	DribblesSuccPercent     string `json:"dribbles_success_percent"`
	FailedGainBallControl   string `json:"failed_gain_ball_control"`
	LooseBallAfterTackle    string `json:"loose_ball_after_tackled"`
	ReceivedPass            string `json:"received_ball"`
	PrgPassReceived         string `json:"progressive_pass_received"`
}

type PlayingTimeStats struct{}

type MiscellaneousStats struct {
	YellowCards           string `json:"yellow_cards"`
	RedCards              string `json:"red_cards"`
	TwoYellowCards        string `json:"two_yellow_cards"`
	Fouls                 string `json:"fouls"`
	Offsides              string `json:"offside"`
	Crosses               string `json:"crosses"`
	Interceptions         string `json:"interceptions"`
	TacklesWonPossesion   string `json:"tackles_won_possession"`
	PenaltyKicksWon       string `json:"penalty_kicks_won"`
	PenaltyKicksConceded  string `json:"penalty_kicks_conceded"`
	OwnGoals              string `json:"own_goals"`
	LooseBallRecovered    string `json:"loose_ball_recov"`
	AerialDuelsWon        string `json:"aerial_duels_won"`
	AerialDuelsLost       string `json:"aerial_duels_lost"`
	AerialDuelsWonPercent string `json:"aerial_duels_won_percent"`
}

func NewMiscellaneousStats(
	yellowCards,
	redCards,
	twoYellowCards,
	fouls,
	offsides,
	crosses,
	interceptions,
	tacklesWonPossesion,
	penaltyKicksWon,
	penaltyKicksConceded,
	ownGoals,
	looseBallRecovered,
	aerialDuelsWon,
	aerialDuelsLost,
	aerialDuelsWonPercent string,
) *MiscellaneousStats {
	return &MiscellaneousStats{
		YellowCards:           yellowCards,
		RedCards:              redCards,
		TwoYellowCards:        twoYellowCards,
		Fouls:                 fouls,
		Offsides:              offsides,
		Crosses:               crosses,
		Interceptions:         interceptions,
		TacklesWonPossesion:   tacklesWonPossesion,
		PenaltyKicksWon:       penaltyKicksWon,
		PenaltyKicksConceded:  penaltyKicksConceded,
		OwnGoals:              ownGoals,
		LooseBallRecovered:    looseBallRecovered,
		AerialDuelsWon:        aerialDuelsWon,
		AerialDuelsLost:       aerialDuelsLost,
		AerialDuelsWonPercent: aerialDuelsWonPercent,
	}
}

type PlayerSummaryStats struct {
	PlayedMatches string `json:"played_matches"`
	PlayedMin     string `json:"played_minutes"`
	Gls           string `json:"goals"`
	Assists       string `json:"assists"`
}

func NewPlayerSummary(playedMatches, playedMin, gls, assists string) *PlayerSummaryStats {
	return &PlayerSummaryStats{
		PlayedMatches: playedMatches,
		PlayedMin:     playedMin,
		Gls:           gls,
		Assists:       assists,
	}
}

type GoalKeepingStats struct {
	GoalAgainst                   string `json:"goal_against"`
	PenaltyKickAllowed            string `json:"pk_allowed"`
	FreeKickAllowed               string `json:"fk_allowed"`
	CornerGoalsAgainst            string `json:"ck_allowed"`
	OwnGoals                      string `json:"own_goals"`
	PostShotxG                    string `json:"ps_xG"`
	PostShotxGSot                 string `json:"ps_xG_SoT"`
	PostShotxGMinusGlsAllowed     string `json:"ps_xG_minus_gls_allowed"`
	PostShotxGMinusGlsAllowedFt   string `json:"ps_xG_minus_gls_allowed_ft"`
	LaunchedPassesCompletedLonger string `json:"passes_cmp_longer"`
	LaunchedPassesAttemptedLonger string `json:"passes_att_longer"`
	LaunchedPassesCompletedPer    string `json:"passes_cmp_per"`
	PassesAttempted               string `json:"passes_att"`
	ThrowsAttempted               string `json:"thr_att"`
	LaunchPer                     string `json:"launch_per"`
	AverageLen                    string `json:"avg_len"`
	GkAttempted                   string `json:"gk_att"`
	GkLaunchPer                   string `json:"gk_launch_per"`
	GkAverageLen                  string `json:"gk_avg_len"`
	OpponentAttemptedCrosses      string `json:"crosses_opp"`
	CrossesStopedByKeeper         string `json:"crosses_stp_by_keeper"`
	CrossesStopedByKeeperPer      string `json:"crosses_stp_by_keeper_per"`
	NDefensiveActionsByKeeper     string `json:"n_defense_actions_by_keeper"`
	NDefensiveActionsByKeeperFt   string `json:"n_defense_actions_by_keeper_ft"`
	AverageDistance               string `json:"average_distance"`
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
	Min    string `json:"minutes_played_per90"`

	Shooting         ShootingStats         `json:"shooting_stats"`
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

func (p *Player) AppendPassingStats(pss PassingStats) {
	p.Passing = pss
}

func (p *Player) AppendPassTypes(pss PassTypesStats) {
	p.PassType = pss
}

func (p *Player) AppendGoalCreation(gsa GoalShotCreationStats) {
	p.GoalShotCreation = gsa
}

func (p *Player) AppendDefensiveActions(df DefenseStats) {
	p.DefenseActions = df
}

func (p *Player) AppendPossession(po PossessionStats) {
	p.Possession = po
}

func (p *Player) AppendPlayingTime(pt PlayingTimeStats) {
	p.PlayingTime = pt
}

func (p *Player) AppendMisscellaneous(mis MiscellaneousStats) {
	p.Miscellaneous = mis
}

func (p *Player) AppendPlayerSummary(ps PlayerSummaryStats) {
	p.PlayerSummary = ps
}
