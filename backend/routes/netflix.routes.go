package routes

import (
	"backend/controllers"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()
	Router.HandleFunc("/", controllers.Home).Methods("GET")
	Router.HandleFunc("/add", controllers.AddMovie).Methods("POST")
	Router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
}
