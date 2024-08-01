package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Metrics struct {
	SessionID        string            `json:"session_id"`
	PlayerID         string            `json:"player_id"`
	StartTime        time.Time         `json:"start_time"`
	EndTime          time.Time         `json:"end_time"`
	FinalLength      int               `json:"final_length"`
	TimeToLength     []LengthTime      `json:"time_to_length"`
	MeanTimeToFruit  float64           `json:"mean_time_to_fruit"`
	DirectionChanges []DirectionChange `json:"direction_changes"`
	InputsToFruit    []InputsToFruit   `json:"inputs_to_fruit"`
	PathFitness      []PathFitness     `json:"path_fitness"`
	Heatmap          []Heatmap         `json:"heatmap"`
	GameOver         GameOver          `json:"game_over"`
}

type LengthTime struct {
	Length    int       `json:"length"`
	TimeSince float64   `json:"time_since"`
	Timestamp time.Time `json:"timestamp"`
}

type DirectionChange struct {
	Direction string    `json:"direction"`
	Timestamp time.Time `json:"timestamp"`
}

type InputsToFruit struct {
	FruitNumber int `json:"fruit_number"`
	Inputs      int `json:"inputs"`
}

type PathFitness struct {
	FruitNumber int     `json:"fruit_number"`
	ActualPath  int     `json:"actual_path"`
	OptimalPath int     `json:"optimal_path"`
	PathRatio   float32 `json:"path_ratio"`
}

type Heatmap struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Visits int `json:"visits"`
}

type GameOver struct {
	Cause    string `json:"cause"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
}

func submitData(w http.ResponseWriter, r *http.Request) {
	var data Metrics
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process and store data (example with Firestore)
	// ctx := context.Background()
	// client, err := firestore.NewClient(ctx, "your-project-id")
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	// defer client.Close()

	// _, _, err = client.Collection("game_data").Add(ctx, data)
	// if err != nil {
	// 	log.Fatalf("Failed to add data: %v", err)
	// }

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Data received successfully")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/submit", submitData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
