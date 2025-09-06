
package models

import "time"

// Models
type Product struct {
	ID          string  `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Category    string  `json:"category" bson:"category"`
	ImageURL    string  `json:"image_url" bson:"image_url"`
	InStock     bool    `json:"in_stock" bson:"in_stock"`
	Weight      string  `json:"weight" bson:"weight"`
}

type CartItem struct {
	ProductID string `json:"product_id" bson:"product_id"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}

type CustomerInfo struct {
	Name    string `json:"name" bson:"name"`
	Phone   string `json:"phone" bson:"phone"`
	Email   string `json:"email" bson:"email"`
	Address string `json:"address" bson:"address"`
}

type Order struct {
	ID           string       `json:"id" bson:"id"`
	CustomerInfo CustomerInfo `json:"customer_info" bson:"customer_info"`
	Items        []CartItem   `json:"items" bson:"items"`
	TotalAmount  float64      `json:"total_amount" bson:"total_amount"`
	Status       string       `json:"status" bson:"status"`
	OrderDate    time.Time    `json:"order_date" bson:"order_date"`
	Notes        string       `json:"notes,omitempty" bson:"notes,omitempty"`
}
