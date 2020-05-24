package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Roll model
type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

var rolls []Roll

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)

	rolls = append(rolls, newRoll)

	json.NewEncoder(w).Encode(newRoll)
}

func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)

			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]

			rolls = append(rolls, newRoll)

			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(rolls)
}

func main() {
	// Generate mockup data
	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"})

	// Initialize router
	router := mux.NewRouter()

	// Endpoint
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
	fmt.Println("Server is listening on localhost:5000")
}

//
