package handlers 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"cmd/pkg"
	"strconv"

	"github.com/gorilla/mux"
)

func StartServer() {
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/tennis_players", Tennis_players).Methods("GET")
	router.HandleFunc("/tennis_players/{number}", PlayersInfo).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8000", router)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}



func Tennis_players(w http.ResponseWriter, r *http.Request) {
	log.Println("entering tennis_players end point")
	var response pkg.Response
	persons := pkg.PrepareResponse()

	response.Tennis_players = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func PlayersInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("entering book info end point")
	vars := mux.Vars(r)
	playerId, err := strconv.Atoi(vars["number"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid book number")
		return
	}

	players := pkg.PrepareResponse()

	var foundPlayer *pkg.Tennis_player
	for _, p := range players {
		if p.Id == playerId {
			foundPlayer = &p
			break
		}
	}

	if foundPlayer == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Book with number %d not found", playerId)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(foundPlayer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling book info")
		return
	}

	w.Write(jsonResponse)
}