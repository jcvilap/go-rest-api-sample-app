package route

import (
	"net/http"
	"github.com/gorilla/mux"
	"controller"
)

func New() http.Handler {
	return routes()
}

func routes() *mux.Router {
	router := mux.NewRouter()
	
	// App Routes
	router.HandleFunc("/", MainController).Methods("GET")
	
	return router
}