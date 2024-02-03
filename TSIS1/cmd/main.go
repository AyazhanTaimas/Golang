package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/AyazhanTaimas/Golang/TSIS1/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/tennis_players", Tennis_players).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8000", router)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

