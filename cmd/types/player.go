package types

type ShootingStats struct {
	Gls         string `json:"Gls"`
	Sht         string `json:"Sh"`
	Sot         string `json:"SoT"`
	Sot_per     string `json:"SoT%"`
	Sht_ft      string `json:"Sh/90"`
	Sot_ft      string `json:"SoT/90"`
	Gls_shot    string `json:"G/Sh"`
	Gls_per_sot string `json:"G/SoT"`
	Dist        string `json:"Dist"`
	Fk          string `json:"FK"`
	Pk          string `json:"PK"`
	PkAtt       string `json:"PKatt"`
	Xg          string `json:"xG"`
	Npxg        string `json:"npxG"`
	Npxg_shot   string `json:"npxG/Sh"`
	Gls_xg      string `json:"G-xG"`
	Np          string `json:"np:G-xG"`
}

func NewShootingStats(
	gls,
	sht,
	sot,
	sot_per,
	sht_ft,
	sot_ft,
	gls_shot,
	gls_per_sot,
	dist,
	fk,
	pk,
	pkAtt,
	xg,
	npxg,
	npxg_shot,
	gls_xg,
	np string,
) *ShootingStats {
	return &ShootingStats{
		Gls:         gls,
		Sht:         sht,
		Sot:         sot,
		Sot_per:     sot_per,
		Sht_ft:      sht_ft,
		Sot_ft:      sot_ft,
		Gls_shot:    gls_shot,
		Gls_per_sot: gls_per_sot,
		Dist:        dist,
		Fk:          fk,
		Pk:          pk,
		PkAtt:       pkAtt,
		Xg:          xg,
		Npxg:        npxg,
		Npxg_shot:   npxg_shot,
		Gls_xg:      gls_xg,
		Np:          np,
	}
}

type PassingStats struct {
	PassesCompleted                  string `json:"Cmp"`
	PassesAttempted                  string `json:"Att"`
	CompletionPercent                string `json:"Cmp%"`
	TotalDist                        string `json:"TotDist"`
	PrgDist                          string `json:"PrgDist"`
	ShortPassesCompleted             string `json:"Cmp_short"`
	ShortPassesAttempted             string `json:"Att_short"`
	ShortPassesCmpPercent            string `json:"Cmp%_short"`
	MediumPassesCompleted            string `json:"Cmp_medium"`
	MediumPassesAttempted            string `json:"Att_medium"`
	MediumPassesCmpPercent           string `json:"Cmp%_medium"`
	LongPassesCompleted              string `json:"Cmp_long"`
	LongPassesAttempted              string `json:"Att_long"`
	LongPassesCmpPercent             string `json:"Comp%_long"`
	Assists                          string `json:"Ast"`
	XAG                              string `json:"xAG"`
	XA                               string `json:"xA"`
	AssistsMinusXAG                  string `json:"A-xAG"`
	AssistedShot                     string `json:"KP"`
	CompletedPassesThatEnterOneThird string `json:"1/3"`
	CompletedPassesIntoTheBox        string `json:"PPA"`
	CompletedCrossesIntoTheBox       string `json:"CrsPA"`
	PrgPasses                        string `json:"Prog"`
}

func NewPassingStats(
	passesCompleted,
	passesAttempted,
	completionPercent,
	totalDist,
	prgDist,
	shortPassesCompleted,
	shortPassesAttempted,
	shortPassesCmpPercent,
	mediumPassesCompleted,
	mediumPassesAttempted,
	mediumPassesCmpPercent,
	longPassesCompleted,
	longPassesAttempted,
	longPassesCmpPercent,
	assists,
	xAG,
	xA,
	assistsMinusXAG,
	assistedShot,
	completedPassesThatEnterOneThird,
	completedPassesIntoTheBox,
	completedCrossesIntoTheBox,
	prgPasses string,
) *PassingStats {
	return &PassingStats{
		PassesCompleted:                  passesCompleted,
		PassesAttempted:                  passesAttempted,
		CompletionPercent:                completionPercent,
		TotalDist:                        totalDist,
		PrgDist:                          prgDist,
		ShortPassesCompleted:             shortPassesCompleted,
		ShortPassesAttempted:             shortPassesAttempted,
		ShortPassesCmpPercent:            shortPassesCmpPercent,
		MediumPassesCompleted:            mediumPassesCompleted,
		MediumPassesAttempted:            mediumPassesAttempted,
		MediumPassesCmpPercent:           mediumPassesCmpPercent,
		LongPassesCompleted:              longPassesCompleted,
		LongPassesAttempted:              longPassesAttempted,
		LongPassesCmpPercent:             longPassesCmpPercent,
		Assists:                          assists,
		XAG:                              xAG,
		XA:                               xA,
		AssistsMinusXAG:                  assistsMinusXAG,
		AssistedShot:                     assistedShot,
		CompletedPassesThatEnterOneThird: completedPassesThatEnterOneThird,
		CompletedPassesIntoTheBox:        completedPassesIntoTheBox,
		CompletedCrossesIntoTheBox:       completedCrossesIntoTheBox,
		PrgPasses:                        prgPasses,
	}
}

type PassTypesStats struct {
	PassesAttempted                 string `json:"Att"`
	LiveBallPasses                  string `json:"Live"`
	DeadBallPasses                  string `json:"Dead"`
	PassesAttemptedFromFreeKicks    string `json:"FK"`
	CompletedPassesIntoOpenSpace    string `json:"TB"`
	PassesThatTravelMoreThan40yards string `json:"Sw"`
	Crosses                         string `json:"Crs"`
	ThrowInTaken                    string `json:"TI"`
	CornerKicks                     string `json:"CK"`
	InSwingingCornerKicks           string `json:"In_ck"`
	OutSwingingCornerKicks          string `json:"Out_ck"`
	StraightCornerKicks             string `json:"Str_ck"`
	OutComesPassesCompleted         string `json:"Cmp_outcome"`
	OffsidePasses                   string `json:"Off_outcome"`
	PassesBlockedByOpp              string `json:"Blocks_outcome"`
}

func NewPassTypes(
	passesAttempted,
	liveBallPasses,
	deadBallPasses,
	passesAttemptedFromFreeKicks,
	completedPassesIntoOpenSpace,
	passesThatTravelMoreThan40yards,
	crosses,
	throwInTaken,
	cornerKicks,
	inSwingingCornerKicks,
	outSwingingCornerKicks,
	straightCornerKicks,
	outComesPassesCompleted,
	offsidePasses,
	passesBlockedByOpp string,
) *PassTypesStats {
	return &PassTypesStats{
		PassesAttempted:                 passesAttempted,
		LiveBallPasses:                  liveBallPasses,
		DeadBallPasses:                  deadBallPasses,
		PassesAttemptedFromFreeKicks:    passesAttemptedFromFreeKicks,
		CompletedPassesIntoOpenSpace:    completedPassesIntoOpenSpace,
		PassesThatTravelMoreThan40yards: passesThatTravelMoreThan40yards,
		Crosses:                         crosses,
		ThrowInTaken:                    throwInTaken,
		CornerKicks:                     cornerKicks,
		InSwingingCornerKicks:           inSwingingCornerKicks,
		OutSwingingCornerKicks:          outSwingingCornerKicks,
		StraightCornerKicks:             straightCornerKicks,
		OutComesPassesCompleted:         outComesPassesCompleted,
		OffsidePasses:                   offsidePasses,
		PassesBlockedByOpp:              passesBlockedByOpp,
	}
}

type GoalShotCreationStats struct {
	ShotCreatingActions              string `json:"SCA"`
	ShotCreatingActionsPer90         string `json:"SCA90"`
	SCPassLiveBall                   string `json:"PassLive_sca"`
	SCPassDeadBall                   string `json:"PassDead_sca"`
	SCSuccessDribleLeadShot          string `json:"Drib_sca"`
	SCShotThatLeadAnotherShot        string `json:"Sh_sca"`
	SCFoulsDrawnThatLeadToShot       string `json:"Fld_sca"`
	SCDefensiveActionsThatLeadToShot string `json:"Def_sca"`
	GoalCreationAction               string `json:"GCA"`
	GoalCreationActionPer90          string `json:"GCA90"`
	GCPassLiveBall                   string `json:"PassLive_gca"`
	GCPassDeadBall                   string `json:"PassDead_gca"`
	GCSuccessDribleLeadShot          string `json:"Drib_gca"`
	GCShotThatLeadAnotherShot        string `json:"Sh_gca"`
	GCFoulsDrawnThatLeadToShot       string `json:"Fld_gca"`
	GCDefensiveActionsThatLeadToShot string `json:"Def_gca"`
}

func NewGoalShotCreation(
	shotCreatingActions,
	shotCreatingActionsPer90,
	sCPassLiveBall,
	sCPassDeadBall,
	sCSuccessDribleLeadShot,
	sCShotThatLeadAnotherShot,
	sCFoulsDrawnThatLeadToShot,
	sCDefensiveActionsThatLeadToShot,
	goalCreationAction,
	goalCreationActionPer90,
	gCPassLiveBall,
	gCPassDeadBall,
	gCSuccessDribleLeadShot,
	gCShotThatLeadAnotherShot,
	gCFoulsDrawnThatLeadToShot,
	gCDefensiveActionsThatLeadToShot string,
) *GoalShotCreationStats {
	return &GoalShotCreationStats{
		ShotCreatingActions:              shotCreatingActions,
		ShotCreatingActionsPer90:         shotCreatingActionsPer90,
		SCPassLiveBall:                   sCPassLiveBall,
		SCPassDeadBall:                   sCPassDeadBall,
		SCSuccessDribleLeadShot:          sCSuccessDribleLeadShot,
		SCShotThatLeadAnotherShot:        sCShotThatLeadAnotherShot,
		SCFoulsDrawnThatLeadToShot:       sCFoulsDrawnThatLeadToShot,
		SCDefensiveActionsThatLeadToShot: sCDefensiveActionsThatLeadToShot,
		GoalCreationAction:               goalCreationAction,
		GoalCreationActionPer90:          goalCreationActionPer90,
		GCPassLiveBall:                   gCPassLiveBall,
		GCPassDeadBall:                   gCPassDeadBall,
		GCSuccessDribleLeadShot:          gCSuccessDribleLeadShot,
		GCShotThatLeadAnotherShot:        gCShotThatLeadAnotherShot,
		GCFoulsDrawnThatLeadToShot:       gCFoulsDrawnThatLeadToShot,
		GCDefensiveActionsThatLeadToShot: gCDefensiveActionsThatLeadToShot,
	}
}

type DefenseStats struct {
	Tackles                   string `json:"Tkl"`
	TacklesWonPossession      string `json:"TklW"`
	TacklesDefensive3rd       string `json:"Def3rd"`
	TacklesMid3rd             string `json:"Mid3rd"`
	TacklesAttacking3rd       string `json:"Att3rd"`
	NumberOfDribblersTackled  string `json:"Tkl_vsDribbles"`
	NumberOfTimesDribblerPast string `json:"Att_vsDribbles"`
	PercentOfDribblersTackled string `json:"Tkl%_vsDribbles"`
	Past                      string `json:"Past"`
	Blocks                    string `json:"Blocks"`
	BlockedShot               string `json:"Blocked_shot"`
	BlockkedPass              string `json:"Blocked_pass"`
	Interceptions             string `json:"Interception"`
	TacklePlusInterception    string `json:"Tkl+Int"`
	Clearances                string `json:"Clearances"`
	ErrorLeadShot             string `json:"Err"`
}

func NewDefenseStats(
	tackles,
	tacklesWonPossession,
	tacklesDefensive3rd,
	tacklesMid3rd,
	tacklesAttacking3rd,
	numberOfDribblersTackled,
	numberOfTimesDribblerPast,
	percentOfDribblersTackled,
	past,
	blocks,
	blockedShot,
	blockkedPass,
	interceptions,
	tacklePlusInterception,
	clearances,
	errorLeadShot string,
) *DefenseStats {
	return &DefenseStats{
		Tackles:                   tackles,
		TacklesWonPossession:      tacklesWonPossession,
		TacklesDefensive3rd:       tacklesDefensive3rd,
		TacklesMid3rd:             tacklesMid3rd,
		TacklesAttacking3rd:       tacklesAttacking3rd,
		NumberOfDribblersTackled:  numberOfDribblersTackled,
		NumberOfTimesDribblerPast: numberOfTimesDribblerPast,
		PercentOfDribblersTackled: percentOfDribblersTackled,
		Past:                      past,
		Blocks:                    blocks,
		BlockedShot:               blockedShot,
		BlockkedPass:              blockkedPass,
		Interceptions:             interceptions,
		TacklePlusInterception:    tacklePlusInterception,
		Clearances:                clearances,
		ErrorLeadShot:             errorLeadShot,
	}
}

type PossessionStats struct {
	Touches                 string `json:"Touches"`
	TouchesDefBoxArea       string `json:"DefPen"`
	TouchesDef3rd           string `json:"Def3rd"`
	TouchesMid3rd           string `json:"Mid3rd"`
	TouchesAttacking3rd     string `json:"Att3rd"`
	TouchesAttackingBoxArea string `json:"AttPen"`
	LiveBallTouches         string `json:"Live"`
	DribblesSucc            string `json:"Succ_dribbles"`
	DribblesAttempted       string `json:"Att_dribbles"`
	DribblesSuccPercent     string `json:"Succ%_dribbles"`
	FailedGainBallControl   string `json:"Mis_dribbles"`
	LooseBallAfterTackle    string `json:"Dis_dribbles"`
	ReceivedPass            string `json:"Rec"`
	PrgPassReceived         string `json:"Prog"`
}

func NewPossessionStats(
	touches,
	touchesDefBoxArea,
	touchesDef3rd,
	touchesMid3rd,
	touchesAttacking3rd,
	touchesAttackingBoxArea,
	liveBallTouches,
	dribblesSucc,
	dribblesAttempted,
	dribblesSuccPercent,
	failedGainBallControl,
	looseBallAfterTackle,
	receivedPass,
	prgPassReceived string,
) *PossessionStats {
	return &PossessionStats{
		Touches:                 touches,
		TouchesDefBoxArea:       touchesDefBoxArea,
		TouchesDef3rd:           touchesDef3rd,
		TouchesMid3rd:           touchesMid3rd,
		TouchesAttacking3rd:     touchesAttacking3rd,
		TouchesAttackingBoxArea: touchesAttackingBoxArea,
		LiveBallTouches:         liveBallTouches,
		DribblesSucc:            dribblesSucc,
		DribblesAttempted:       dribblesAttempted,
		DribblesSuccPercent:     dribblesSuccPercent,
		FailedGainBallControl:   failedGainBallControl,
		LooseBallAfterTackle:    looseBallAfterTackle,
		ReceivedPass:            receivedPass,
		PrgPassReceived:         prgPassReceived,
	}
}

type PlayingTimeStats struct {
	PlayedMatches           string `json:"MP"`
	PlayedMin               string `json:"Min"`
	MinuterPerMatches       string `json:"Mn/MP"`
	MinPercent              string `json:"Min%"`
	MinutesDividedBy90      string `json:"90s"`
	MatchesStarted          string `json:"Starts"`
	MinutesPerMatchStarted  string `json:"Mn/Start"`
	FullTimePlayed          string `json:"Compl"`
	GamesAsSub              string `json:"Subs"`
	MinutesPerSubstitution  string `json:"Mn/Sub"`
	UnsedSub                string `json:"unSub"`
	PointsPerMatch          string `json:"PPM"`
	GlsScoredByTeamOnGame   string `json:"onG"`
	GlsAllowedByTeamOnGame  string `json:"onGA"`
	GlsScoredMinusAllowed   string `json:"+/-"`
	GlsScoredMinusAllowed90 string `json:"+/-90"`
	GlsNetPer90             string `json:"On-Off"`
	OnGamexG                string `json:"onxG"`
	OnGamexGA               string `json:"onxGA"`
	OnGamexGMinusxGA        string `json:"xG+/-"`
	OnGamexGMinusxGA90      string `json:"onxGA+/-90"`
	XGMinusxGANetPer90      string `json:"On-Off_teamSuccess(xG)"`
}

func NewPlayingTimeStats(
	playedMatches,
	playedMin,
	minuterPerMatches,
	minPercent,
	minutesDividedBy90,
	matchesStarted,
	minutesPerMatchStarted,
	fullTimePlayed,
	gamesAsSub,
	minutesPerSubstitution,
	unsedSub,
	pointsPerMatch,
	glsScoredByTeamOnGame,
	glsAllowedByTeamOnGame,
	glsScoredMinusAllowed,
	glsScoredMinusAllowed90,
	glsNetPer90,
	onGamexG,
	onGamexGA,
	onGamexGMinusxGA,
	onGamexGMinusxGA90,
	xGMinusxGANetPer90 string,
) *PlayingTimeStats {
	return &PlayingTimeStats{
		PlayedMatches:           playedMatches,
		PlayedMin:               playedMin,
		MinuterPerMatches:       minuterPerMatches,
		MinPercent:              minPercent,
		MinutesDividedBy90:      minutesDividedBy90,
		MatchesStarted:          matchesStarted,
		MinutesPerMatchStarted:  minutesPerMatchStarted,
		FullTimePlayed:          fullTimePlayed,
		GamesAsSub:              gamesAsSub,
		MinutesPerSubstitution:  minutesPerSubstitution,
		UnsedSub:                unsedSub,
		PointsPerMatch:          pointsPerMatch,
		GlsScoredByTeamOnGame:   glsScoredByTeamOnGame,
		GlsAllowedByTeamOnGame:  glsAllowedByTeamOnGame,
		GlsScoredMinusAllowed:   glsScoredMinusAllowed,
		GlsScoredMinusAllowed90: glsScoredMinusAllowed90,
		GlsNetPer90:             glsNetPer90,
		OnGamexG:                onGamexG,
		OnGamexGA:               onGamexGA,
		OnGamexGMinusxGA:        onGamexGMinusxGA,
		OnGamexGMinusxGA90:      onGamexGMinusxGA90,
		XGMinusxGANetPer90:      xGMinusxGANetPer90,
	}
}

type MiscellaneousStats struct {
	YellowCards           string `json:"CrdY"`
	RedCards              string `json:"CrdR"`
	TwoYellowCards        string `json:"2CrdY"`
	Fouls                 string `json:"Fls"`
	FoulsDrawn            string `json:"Fld"`
	Offsides              string `json:"Off"`
	Crosses               string `json:"Crs"`
	Interceptions         string `json:"Interceptions"`
	TacklesWonPossesion   string `json:"TklW"`
	PenaltyKicksWon       string `json:"PKwon"`
	PenaltyKicksConceded  string `json:"PKcon"`
	OwnGoals              string `json:"OG"`
	LooseBallRecovered    string `json:"Recov"`
	AerialDuelsWon        string `json:"Won"`
	AerialDuelsLost       string `json:"Lost"`
	AerialDuelsWonPercent string `json:"Won%"`
}

func NewMiscellaneousStats(
	yellowCards,
	redCards,
	twoYellowCards,
	fouls,
	foulsDrawn,
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
		FoulsDrawn:            foulsDrawn,
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
	GoalAgainst                   string `json:"GA"`
	PenaltyKickAllowed            string `json:"PKA"`
	FreeKickAllowed               string `json:"FK"`
	CornerGoalsAgainst            string `json:"CK"`
	OwnGoals                      string `json:"OG"`
	PostShotxG                    string `json:"PSxG"`
	PostShotxGSot                 string `json:"PSxG/SoT"`
	PostShotxGMinusGlsAllowed     string `json:"PSxG+/-"`
	PostShotxGMinusGlsAllowedFt   string `json:"/90"`
	LaunchedPassesCompletedLonger string `json:"Cmp"`
	LaunchedPassesAttemptedLonger string `json:"Att"`
	LaunchedPassesCompletedPer    string `json:"Cmp%"`
	PassesAttempted               string `json:"Att_passes"`
	ThrowsAttempted               string `json:"Thr"`
	LaunchPer                     string `json:"Launch%"`
	AverageLen                    string `json:"AvgLen"`
	GkAttempted                   string `json:"Att_gk"`
	GkLaunchPer                   string `json:"Launch%_gk"`
	GkAverageLen                  string `json:"AvgLen_gk"`
	OpponentAttemptedCrosses      string `json:"Opp"`
	CrossesStopedByKeeper         string `json:"Stp"`
	CrossesStopedByKeeperPer      string `json:"Stp%"`
	NDefensiveActionsByKeeper     string `json:"#OPA%"`
	NDefensiveActionsByKeeperFt   string `json:"#OPA/90"`
	AverageDistance               string `json:"AvgDist"`
}

func NewGoalKeepingStats(
	goalAgainst,
	penaltyKickAllowed,
	freeKickAllowed,
	cornerGoalsAgainst,
	ownGoals,
	postShotxG,
	postShotxGSot,
	postShotxGMinusGlsAllowed,
	postShotxGMinusGlsAllowedFt,
	launchedPassesCompletedLonger,
	launchedPassesAttemptedLonger,
	launchedPassesCompletedPer,
	passesAttempted,
	throwsAttempted,
	launchPer,
	averageLen,
	gkAttempted,
	gkLaunchPer,
	gkAverageLen,
	opponentAttemptedCrosses,
	crossesStopedByKeeper,
	crossesStopedByKeeperPer,
	nDefensiveActionsByKeeper,
	nDefensiveActionsByKeeperFt,
	averageDistance string,
) *GoalKeepingStats {
	return &GoalKeepingStats{
		GoalAgainst:                   goalAgainst,
		PenaltyKickAllowed:            penaltyKickAllowed,
		FreeKickAllowed:               freeKickAllowed,
		CornerGoalsAgainst:            cornerGoalsAgainst,
		OwnGoals:                      ownGoals,
		PostShotxG:                    postShotxG,
		PostShotxGSot:                 postShotxGSot,
		PostShotxGMinusGlsAllowed:     postShotxGMinusGlsAllowed,
		PostShotxGMinusGlsAllowedFt:   postShotxGMinusGlsAllowedFt,
		LaunchedPassesCompletedLonger: launchedPassesCompletedLonger,
		LaunchedPassesAttemptedLonger: launchedPassesAttemptedLonger,
		LaunchedPassesCompletedPer:    launchedPassesCompletedPer,
		PassesAttempted:               passesAttempted,
		ThrowsAttempted:               throwsAttempted,
		LaunchPer:                     launchPer,
		AverageLen:                    averageLen,
		GkAttempted:                   gkAttempted,
		GkLaunchPer:                   gkLaunchPer,
		GkAverageLen:                  gkAverageLen,
		OpponentAttemptedCrosses:      opponentAttemptedCrosses,
		CrossesStopedByKeeper:         crossesStopedByKeeper,
		CrossesStopedByKeeperPer:      crossesStopedByKeeperPer,
		NDefensiveActionsByKeeper:     nDefensiveActionsByKeeper,
		NDefensiveActionsByKeeperFt:   nDefensiveActionsByKeeperFt,
		AverageDistance:               averageDistance,
	}
}

type PlayerBasic struct {
	Name   string
	Nation string
	Pos    string
	Age    string
	Min    string
}

func NewPlayerBasic(name, nation, pos, age, min string) *PlayerBasic {
	return &PlayerBasic{
		Name:   name,
		Nation: nation,
		Pos:    pos,
		Age:    age,
		Min:    min,
	}
}

type Player struct {
	Name   string `json:"Name"`
	Nation string `json:"Nation"`
	Pos    string `json:"Position"`
	Age    string `json:"Age"`
	Min    string `json:"Min/90"`

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

type Players struct {
	Team        string   `json:"team"`
	PlayersList []Player `json:"players_stats"`
}

func NewPlayers(team string, pl []Player) *Players {
	return &Players{
		Team:        team,
		PlayersList: pl,
	}
}
