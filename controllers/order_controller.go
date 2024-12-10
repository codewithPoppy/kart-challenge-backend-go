package controllers

import (
	"encoding/json"
	"net/http"

	"kart-challenge-backend/models"
	"kart-challenge-backend/services"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming request body into an OrderReq struct
	var orderReq models.OrderReq
	err := json.NewDecoder(r.Body).Decode(&orderReq)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Use the service to send the order request to the external API
	order, err := services.PlaceOrder(orderReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Respond with the order details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
