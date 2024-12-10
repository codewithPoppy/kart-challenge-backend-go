package controllers

import (
	"encoding/json"
	"net/http"

	"kart-challenge-backend/services"

	"github.com/gorilla/mux"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch products from the external API
	products, err := services.ListProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// Get product ID from the URL path
	vars := mux.Vars(r)
	productId := vars["productId"]

	// Fetch the product by ID from the external API
	product, err := services.GetProductByID(productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
