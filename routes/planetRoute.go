package routes

import (
	"americanas/middleware"
	"github.com/gorilla/mux"
)

func PlanetRoute(router *mux.Router) {
	router.HandleFunc("/planet", middleware.GetPlanet).Methods("GET")
	router.HandleFunc("/planet", middleware.CreatePlanet).Methods("POST")
	router.HandleFunc("/planet/{id}", middleware.GetBookById).Methods("GET")
}
