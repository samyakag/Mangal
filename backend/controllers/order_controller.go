package controllers

import (
	"mangal-chai-backend/models"
	"mangal-chai-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	Service services.OrderServiceInterface
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var orderData struct {
		CustomerInfo models.CustomerInfo `json:"customer_info"`
		Items        []models.CartItem   `json:"items"`
		Notes        string              `json:"notes"`
	}

	if err := ctx.ShouldBindJSON(&orderData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := c.Service.CreateOrder(orderData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID, "total_amount": order.TotalAmount})
}

func (c *OrderController) GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("order_id")
	order, err := c.Service.GetOrder(orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	ctx.JSON(http.StatusOK, order)
}