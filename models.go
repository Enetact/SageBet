package models

type Match struct {
	HomeTeam     string  `json:"home_team"`
	AwayTeam     string  `json:"away_team"`
	HomeAvgGoals float64 `json:"home_avg_goals"`
	AwayAvgGoals float64 `json:"away_avg_goals"`
	OddsHome     float64 `json:"odds_home"`
	OddsAway     float64 `json:"odds_away"`
	OddsDraw     float64 `json:"odds_draw"`
}

type MatchDataRequest struct {
	Matches []Match `json:"matches"`
}

type Prediction struct {
	MatchID        string  `json:"match_id"`
	WinProbability float64 `json:"win_probability"`
	Confidence     string  `json:"confidence"`
	RecommendedBet string  `json:"recommended_bet"`
	StakePercent   float64 `json:"stake_percent"`
	Rationale      string  `json:"rationale"`
}

type AIPredictionResponse struct {
	Predictions []Prediction `json:"predictions"`
}
