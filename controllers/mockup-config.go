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
	w.Header().Set("Content-Type", "application/json")

	db := helpers.OpenDB()
	defer helpers.CloseDB()

	var mockupConfigs []models.MockupConfig
	err := db.All(&mockupConfigs)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mockupConfigs)
}

// Get mockup config
func mockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := helpers.OpenDB()
	defer helpers.CloseDB()

	var mockupConfig models.MockupConfig
	id, _ := strconv.Atoi(params["id"])

	err := db.One("ID", id, &mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mockupConfig)
}

// Create mockup config
func createMockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := helpers.OpenDB()
	defer helpers.CloseDB()

	var mockupConfig models.MockupConfig
	json.NewDecoder(r.Body).Decode(&mockupConfig)
	// TODO: validate body

	err := db.Save(&mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mockupConfig)
}

// Update mockup config
func updateMockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := helpers.OpenDB()
	defer helpers.CloseDB()

	var mockupConfig models.MockupConfig
	id, _ := strconv.Atoi(params["id"])

	err := db.One("ID", id, &mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(r.Body).Decode(&mockupConfig)
	// TODO: validate body

	err = db.Save(&mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(mockupConfig)
}

// Delete mockup config
func deleteMockupConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db := helpers.OpenDB()
	defer helpers.CloseDB()

	var mockupConfig models.MockupConfig
	id, _ := strconv.Atoi(params["id"])

	err := db.One("ID", id, &mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Drop(&mockupConfig)
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
