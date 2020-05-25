package configs

import (
	"fmt"
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

	fmt.Println("Server is listening on localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
