package main

import (
	"fmt"
	"log"
	"net/http"

	"kart-challenge-backend/config"
	"kart-challenge-backend/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize routes
	router := routes.InitializeRoutes()

	// Start the server
	log.Printf("Server running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), router))
}
