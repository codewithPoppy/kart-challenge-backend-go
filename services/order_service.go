package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"kart-challenge-backend/models"
	"net/http"
)

// API endpoint for placing orders
const ordersAPIURL = "https://orderfoodonline.deno.dev/api/order"

// PlaceOrder sends the order request to the external API after validating the promo code.
func PlaceOrder(orderReq models.OrderReq) (models.Order, error) {
	// Validate the promo code if provided
	if orderReq.CouponCode != "" {
		isValid, err := ValidatePromoCode(orderReq.CouponCode)
		if err != nil || !isValid {
			return models.Order{}, errors.New("invalid promo code")
		}
	}

	// Marshal the order request to JSON
	orderReqBody, err := json.Marshal(orderReq)
	if err != nil {
		return models.Order{}, errors.New("failed to marshal order request")
	}

	// Send HTTP POST request to the API
	resp, err := http.Post(ordersAPIURL, "application/json", bytes.NewBuffer(orderReqBody))
	if err != nil {
		return models.Order{}, errors.New("failed to send order request to the API")
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return models.Order{}, errors.New("API returned an error: " + string(body))
	}

	// Parse the response body
	var order models.Order
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return models.Order{}, errors.New("failed to parse API response")
	}

	return order, nil
}
