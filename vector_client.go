package vector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"autobetsage/models"
)

const baseURL = "http://127.0.0.1:5100"

type queryResponse struct {
	Documents [][]string         `json:"documents"`
	Metadatas [][]map[string]any `json:"metadatas"`
	Ids       [][]string         `json:"ids"`
}

// Store sends text + metadata to the vector DB
func Store(text string, match models.Match) {
	data := map[string]interface{}{
		"text": text,
		"meta": map[string]interface{}{
			"home_team": match.HomeTeam,
			"away_team": match.AwayTeam,
			"odds_home": match.OddsHome,
			"odds_away": match.OddsAway,
			"odds_draw": match.OddsDraw,
		},
	}
	payload, _ := json.Marshal(data)
	resp, err := http.Post(baseURL+"/add", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Vector DB store error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Store Result:", string(body))
}

// Query sends a search phrase and returns top matches
func Query(question string) string {
	data := map[string]interface{}{
		"query": question,
		"n":     3,
	}
	payload, _ := json.Marshal(data)
	resp, err := http.Post(baseURL+"/query", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Vector DB query error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result queryResponse
	json.Unmarshal(body, &result)

	fmt.Println("\nTop Results:")
	for i, doc := range result.Documents {
		fmt.Printf("  [%d] %s\n", i+1, doc[0])
	}
	return ""
}
