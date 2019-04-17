package main

import (
	"jokesapi/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/jokes/", api.AllJokes).Methods("GET")
	router.HandleFunc("/post", api.NameWriter).Methods("POST")
	router.HandleFunc("/", api.RandJoke).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	api.JokeDAO()
	handleRequests()
}
