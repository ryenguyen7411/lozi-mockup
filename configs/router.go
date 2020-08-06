package configs

import (
	"fmt"
	"main/controllers"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// HandleRouter ...
func HandleRouter() {
	// Initialize router
	router := mux.NewRouter()

	for _, item := range controllers.MockupConfigRouter {
		router.HandleFunc(item.URL, item.Handler).Methods(item.Method)
	}

	router.PathPrefix("/").HandlerFunc(controllers.MockupAPIHandler)

	fmt.Println("Server is listening on localhost:5000")
	http.ListenAndServe(
		":5000",
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(router),
	)
}
