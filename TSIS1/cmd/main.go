package main

import (
	"log"
	"cmd/internal/handlers"

)

func main() {
	log.Println("starting API server")
	handlers.StartServer()
}


