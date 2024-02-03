package handlers 

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/AyazhanTaimas/Golang/TSIS1/pkg"
	"github.com/gorilla/mux"
)

func Tennis_players(w http.ResponseWriter, r *http.Request) {
	log.Println("entering tennis_players end point")
	var response Response
	persons := prepareResponse()

	response.Tennis_players = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}