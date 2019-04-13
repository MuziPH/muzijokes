package main

import (
	"fmt"
	"jokesapi/api"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleRequests(id string) {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println(id)
	router.HandleFunc("/", api.RandJoke)
	// router.HandleFunc("/jokes", api.AllJokes).Methods("GET")
	router.HandleFunc("/jokes/", api.AllJokes).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	ranNum := rand.Intn(9)
	strNum := strconv.Itoa(ranNum)
	fmt.Println(strNum)
	api.JokeDAO()
	handleRequests(strNum)
}
