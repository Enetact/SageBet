package ai

import (
	"context"
	"fmt"
	"os"
	"strings"

	"autobetsage/models"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func init() {
	client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}

func EvaluateBets(req models.MatchDataRequest) (*models.AIPredictionResponse, error) {
	var results []models.Prediction

	for _, match := range req.Matches {
		prompt := buildPrompt(match)

		resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "system",
					Content: "You are a sports betting expert. Respond only with valid JSON.",
				},
				{
					Role:    "user",
					Content: prompt,
				},
			},
		})
		if err != nil {
			return nil, err
		}

		// Simulated basic output parsing
		results = append(results, models.Prediction{
			MatchID:        fmt.Sprintf("%s_vs_%s", match.HomeTeam, match.AwayTeam),
			WinProbability: 0.65,
			Confidence:     "High",
			RecommendedBet: "Home Win",
			StakePercent:   4.0,
			Rationale:      strings.TrimSpace(resp.Choices[0].Message.Content),
		})
	}

	return &models.AIPredictionResponse{Predictions: results}, nil
}

func buildPrompt(match models.Match) string {
	return fmt.Sprintf(`Analyze this match:

- Home Team: %s
- Away Team: %s
- Avg Home Goals: %.2f
- Avg Away Goals: %.2f
- Odds: Home %.2f | Draw %.2f | Away %.2f

Provide a JSON output with: win_probability (0-1), confidence, recommended_bet, stake_percent (0-5), and rationale.`,
		match.HomeTeam, match.AwayTeam, match.HomeAvgGoals, match.AwayAvgGoals,
		match.OddsHome, match.OddsDraw, match.OddsAway)
}
