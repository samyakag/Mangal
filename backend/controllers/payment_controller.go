
package controllers

import (
	"net/http"
	"os"

	"mangal-chai-backend/services"

	"github.com/gin-gonic/gin"
)

// PaymentController handles payment related requests
type PaymentController struct {
	Service *services.PaymentService
}

// CreateRazorpayOrder creates a new Razorpay order
func (pc *PaymentController) CreateRazorpayOrder(c *gin.Context) {
	var req services.CreateRazorpayOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := pc.Service.CreateRazorpayOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id": order["id"],
		"amount":   order["amount"],
		"currency": order["currency"],
		"key_id":   os.Getenv("RAZORPAY_KEY_ID"),
	})
}
