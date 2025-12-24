
package services

import (
	"mangal-chai-backend/models"
	"os"

	"github.com/razorpay/razorpay-go"
)

// PaymentService handles payment related logic
type PaymentService struct {
	razorpayClient *razorpay.Client
}

type CreateRazorpayOrderRequest struct {
	Items []models.CartItem `json:"items"`
}

// NewPaymentService creates a new PaymentService
func NewPaymentService() *PaymentService {
	keyId := os.Getenv("RAZORPAY_KEY_ID")
	keySecret := os.Getenv("RAZORPAY_KEY_SECRET")

	if keyId == "" || keySecret == "" {
		panic("RAZORPAY_KEY_ID or RAZORPAY_KEY_SECRET environment variable not set")
	}

	client := razorpay.NewClient(keyId, keySecret)

	return &PaymentService{razorpayClient: client}
}

// CreateRazorpayOrder creates a new Razorpay order
func (ps *PaymentService) CreateRazorpayOrder(request CreateRazorpayOrderRequest) (map[string]interface{}, error) {
	// This is a placeholder for calculating the order amount based on the items
	// In a real application, you would fetch the product prices from your database
	// and calculate the total amount
	var totalAmount int64 = 100000 // Placeholder amount (e.g., 1000.00 INR)

	orderParams := map[string]interface{}{
		"amount":   totalAmount,
		"currency": "INR",
		"receipt":  "some_receipt_id", // Replace with a unique receipt ID
	}

	order, err := ps.razorpayClient.Order.Create(orderParams, nil)
	if err != nil {
		return nil, err
	}

	return order, nil
}
