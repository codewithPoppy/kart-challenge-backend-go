package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kart-challenge-backend/models"
	"net/http"
)

// API endpoint for retrieving products
const productsAPIURL = "https://orderfoodonline.deno.dev/api/product"

// ListProducts retrieves all available products from the external API
func ListProducts() ([]models.Product, error) {
	// Send HTTP GET request to the API
	resp, err := http.Get(productsAPIURL)
	if err != nil {
		return nil, errors.New("failed to fetch products from the API")
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("API returned a non-200 status code")
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read API response")
	}

	// Unmarshal the response into a slice of products
	var products []models.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, errors.New("failed to parse API response")
	}

	return products, nil
}

// GetProductByID retrieves a product by its ID from the external API
func GetProductByID(productId string) (models.Product, error) {
	// Retrieve all products
	products, err := ListProducts()
	if err != nil {
		return models.Product{}, err
	}

	// Search for the product with the given ID
	for _, product := range products {
		if product.ID == productId {
			return product, nil
		}
	}

	return models.Product{}, errors.New("product not found")
}