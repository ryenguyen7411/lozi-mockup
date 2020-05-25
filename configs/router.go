package configs

import (
	"log"
	"main/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleRouter ...
func HandleRouter() {
	// Initialize router
	router := mux.NewRouter()

	for _, item := range controllers.MockupConfigRouter {
		router.HandleFunc(item.URL, item.Handler).Methods(item.Method)
	}

	log.Fatal(http.ListenAndServe(":5000", router))
}
