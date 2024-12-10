package models

// OrderReq represents the request structure for placing an order
type OrderReq struct {
	CouponCode string      `json:"couponCode,omitempty"` // Optional promo code
	Items      []OrderItem `json:"items"`                // List of items in the order
}

// OrderItem represents an individual item in an order
type OrderItem struct {
	ProductId string `json:"productId"` // Product ID (required)
	Quantity  int    `json:"quantity"`  // Quantity of the product (required)
}

// Order represents the response structure for an order
type Order struct {
	ID       string      `json:"id"`
	Items    []OrderItem `json:"items"`
	Products []Product   `json:"products"`
}
