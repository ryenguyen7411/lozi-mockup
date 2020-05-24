package configs

import (
	"log"
	"main/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RouterConfig ...
type RouterConfig struct {
	Method, URL string
	Handler     func(w http.ResponseWriter, r *http.Request)
}

// HandleRouter ...
func HandleRouter() {
	// Initialize router
	router := mux.NewRouter()

	// var data []RouterConfig

	// data = append(data, controllers.MockupConfigRouter...)

	for _, item := range controllers.MockupConfigRouter {
		router.HandleFunc(item.URL, item.Handler).Methods(item.Method)
	}

	log.Fatal(http.ListenAndServe(":5000", router))
}
