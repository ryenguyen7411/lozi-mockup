package controllers

import (
	"net/http"
)

// Get list of mockup config
func mockupConfigs(w http.ResponseWriter, r *http.Request) {

}

// Get mockup config
func mockupConfig(w http.ResponseWriter, r *http.Request) {

}

// Create mockup config
func createMockupConfig(w http.ResponseWriter, r *http.Request) {

}

// Update mockup config
func updateMockupConfig(w http.ResponseWriter, r *http.Request) {

}

// Delete mockup config
func deleteMockupConfig(w http.ResponseWriter, r *http.Request) {

}

// RouterConfig ...
type RouterConfig struct {
	Method, URL string
	Handler     func(w http.ResponseWriter, r *http.Request)
}

// MockupConfigRouter ...
var MockupConfigRouter = []RouterConfig{
	{"GET", "/mockup-configs", mockupConfigs},
	{"GET", "/mockup-config/{id}", mockupConfig},
	{"POST", "/mockup-config", createMockupConfig},
	{"PUT", "/mockup-config/{id}", updateMockupConfig},
	{"DELETE", "/mockup-config/{id}", deleteMockupConfig},
}
