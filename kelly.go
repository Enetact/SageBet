package models

func KellyStake(probability, odds float64) float64 {
	q := 1 - probability
	b := odds - 1
	if b <= 0 {
		return 0
	}
	fraction := ((b * probability) - q) / b
	if fraction < 0 {
		return 0
	}
	return fraction
}
