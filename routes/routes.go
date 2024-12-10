package routes

import (
	"kart-challenge-backend/config"
	"kart-challenge-backend/controllers"
	"kart-challenge-backend/middleware"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	cfg := config.LoadConfig()

	router := mux.NewRouter()
	router.Use(middleware.RequestLogger)              // Log all requests
	router.Use(middleware.ValidateAPIKey(cfg.APIKey)) // Validate API key

	// Products API
	router.HandleFunc("/product", controllers.ListProducts).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.GetProduct).Methods("GET")

	// Orders API
	router.HandleFunc("/order", controllers.PlaceOrder).Methods("POST")

	return router
}
