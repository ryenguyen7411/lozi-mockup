package controllers

import (
	"encoding/json"
	"log"
	"main/helpers"
	"main/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get list of mockup config
func mockupConfigs(w http.ResponseWriter, r *http.Request) {

}

// Get mockup config
func mockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Get mockup config by id
	id, err := strconv.Atoi(params["id"])
	bytes, err := helpers.Read(helpers.Itob(id))
	if err != nil {
		log.Fatal(err)
	}

	var mockupConfig models.MockupConfig
	err = json.Unmarshal([]byte(bytes), &mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mockupConfig)
}

// Create mockup config
func createMockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Validate body
	var newMockupConfig models.MockupConfig
	json.NewDecoder(r.Body).Decode(&newMockupConfig)
	// TODO: validate body

	nextID := helpers.GetNextSequence()
	newMockupConfig.ID = int(nextID)

	// Save to DB
	bytes, err := json.Marshal(newMockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	_, err = helpers.Create(helpers.Itob(newMockupConfig.ID), bytes)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(newMockupConfig)
}

// Update mockup config
func updateMockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Get mockup config by id
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := helpers.Read(helpers.Itob(id))
	if err != nil {
		log.Fatal(err)
	}

	var mockupConfig models.MockupConfig
	err = json.Unmarshal([]byte(bytes), &mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Validate body
	json.NewDecoder(r.Body).Decode(&mockupConfig)
	// TODO: validate body

	// Save to DB
	newBytes, err := json.Marshal(mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	_, err = helpers.Update(helpers.Itob(mockupConfig.ID), newBytes)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mockupConfig)
}

// Delete mockup config
func deleteMockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Delete mockup config by id
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := helpers.Delete(helpers.Itob(id))
	if err != nil {
		log.Fatal(err)
	}

	var mockupConfig models.MockupConfig
	err = json.Unmarshal([]byte(bytes), &mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mockupConfig)
}

// RouterConfig ...
type RouterConfig struct {
	Method, URL string
	Handler     func(w http.ResponseWriter, r *http.Request)
}

// MockupConfigRouter ...
var MockupConfigRouter = []RouterConfig{
	{"GET", "/mockup-configs", mockupConfigs},
	{"GET", "/mockup-configs/{id}", mockupConfig},
	{"POST", "/mockup-configs", createMockupConfig},
	{"PUT", "/mockup-configs/{id}", updateMockupConfig},
	{"DELETE", "/mockup-configs/{id}", deleteMockupConfig},
}
