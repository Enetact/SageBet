package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"autobetsage/ai"
	"autobetsage/models"
)

func main() {
	http.HandleFunc("/ai-picks", handleAIPicks)
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("AutoBetSage.AI is live"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleAIPicks(w http.ResponseWriter, r *http.Request) {
	var input models.MatchDataRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	resp, err := ai.EvaluateBets(input)
	if err != nil {
		http.Error(w, "AI Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
