package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"autobetsage/models"
)

type GameData struct {
	Match  models.Match
	League string
	Sport  string
	Start  time.Time
	Book   string
}

var oddsAPIKey = os.Getenv("ODDS_API_KEY")
var footyAPIKey = os.Getenv("FOOTBALL_API_KEY")

// Pulls odds from TheOddsAPI
func GetTodayOdds(sportKey string) ([]GameData, error) {
	url := fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds/?regions=us&markets=h2h&oddsFormat=decimal&apiKey=%s", sportKey, oddsAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("TheOddsAPI failed: %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var games []map[string]interface{}
	json.Unmarshal(body, &games)

	var results []GameData
	for _, g := range games {
		commence, _ := time.Parse(time.RFC3339, g["commence_time"].(string))
		teams := g["teams"].([]interface{})
		bookmakers := g["bookmakers"].([]interface{})
		if len(bookmakers) == 0 {
			continue
		}
		book := bookmakers[0].(map[string]interface{})
		markets := book["markets"].([]interface{})
		if len(markets) == 0 {
			continue
		}
		outcomes := markets[0].(map[string]interface{})["outcomes"].([]interface{})

		home, away := teams[0].(string), teams[1].(string)
		match := models.Match{
			HomeTeam: home,
			AwayTeam: away,
		}
		for _, o := range outcomes {
			od := o.(map[string]interface{})
			switch od["name"].(string) {
			case home:
				match.OddsHome = od["price"].(float64)
			case away:
				match.OddsAway = od["price"].(float64)
			case "Draw":
				match.OddsDraw = od["price"].(float64)
			}
		}

		results = append(results, GameData{
			Match:  match,
			League: g["sport_title"].(string),
			Sport:  g["sport_key"].(string),
			Start:  commence,
			Book:   book["title"].(string),
		})
	}

	return results, nil
}
