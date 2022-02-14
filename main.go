package main

import (
	"americanas/repository"
	"americanas/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	repository.Db()
	routes.PlanetRoute(r)
	log.Fatal(http.ListenAndServe(":5001", r))
}
