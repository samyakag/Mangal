
package services

import (
	"fmt"
	"mangal-chai-backend/models"
	"mangal-chai-backend/repositories"
	"time"
)

type OrderServiceInterface interface {
	CreateOrder(orderData struct {
		CustomerInfo models.CustomerInfo `json:"customer_info"`
		Items        []models.CartItem   `json:"items"`
		Notes        string              `json:"notes"`
	}) (*models.Order, error)
	GetOrder(id string) (*models.Order, error)
}

type OrderService struct {
	OrderRepository   repositories.OrderRepositoryInterface
	ProductRepository repositories.ProductRepositoryInterface
}

func (s *OrderService) CreateOrder(orderData struct {
	CustomerInfo models.CustomerInfo `json:"customer_info"`
	Items        []models.CartItem   `json:"items"`
	Notes        string              `json:"notes"`
}) (*models.Order, error) {
	totalAmount := 0.0
	for _, item := range orderData.Items {
		product, err := s.ProductRepository.GetProduct(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product %s not found", item.ProductID)
		}
		if !product.InStock {
			return nil, fmt.Errorf("product %s is out of stock", product.Name)
		}
		totalAmount += product.Price * float64(item.Quantity)
	}

	newOrder := models.Order{
		ID:           fmt.Sprintf("ord_%d", time.Now().UnixNano()),
		CustomerInfo: orderData.CustomerInfo,
		Items:        orderData.Items,
		TotalAmount:  totalAmount,
		Status:       "pending",
		OrderDate:    time.Now(),
		Notes:        orderData.Notes,
	}

	err := s.OrderRepository.CreateOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &newOrder, nil
}

func (s *OrderService) GetOrder(id string) (*models.Order, error) {
	return s.OrderRepository.GetOrder(id)
}
