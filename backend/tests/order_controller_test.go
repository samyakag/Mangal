
package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"mangal-chai-backend/controllers"
	"mangal-chai-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(orderData struct {
	CustomerInfo models.CustomerInfo `json:"customer_info"`
	Items        []models.CartItem   `json:"items"`
	Notes        string              `json:"notes"`
}) (*models.Order, error) {
	args := m.Called(orderData)
	val := args.Get(0)
	if val == nil {
		return nil, args.Error(1)
	}
	return val.(*models.Order), args.Error(1)
}

func (m *MockOrderService) GetOrder(id string) (*models.Order, error) {
	args := m.Called(id)
	val := args.Get(0)
	if val == nil {
		return nil, args.Error(1)
	}
	return val.(*models.Order), args.Error(1)
}

func TestOrderController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Test CreateOrder
	t.Run("CreateOrder - Success", func(t *testing.T) {
		mockService := new(MockOrderService)
		expectedOrder := &models.Order{ID: "order1", TotalAmount: 10.0}
		mockService.On("CreateOrder", mock.Anything).Return(expectedOrder, nil)

		controller := &controllers.OrderController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		orderData := map[string]interface{}{
			"customer_info": map[string]string{"name": "John Doe"},
			"items":         []map[string]interface{}{{"product_id": "prod1", "quantity": 1}},
			"notes":         "",
		}
		jsonValue, _ := json.Marshal(orderData)
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/orders", bytes.NewBuffer(jsonValue))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.CreateOrder(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Order placed successfully")
		mockService.AssertExpectations(t)
	})

	t.Run("CreateOrder - Service Error", func(t *testing.T) {
		mockService := new(MockOrderService)
		mockService.On("CreateOrder", mock.Anything).Return(nil, errors.New("service error"))

		controller := &controllers.OrderController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		orderData := map[string]interface{}{
			"customer_info": map[string]string{"name": "John Doe"},
			"items":         []map[string]interface{}{{"product_id": "prod1", "quantity": 1}},
			"notes":         "",
		}
		jsonValue, _ := json.Marshal(orderData)
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/orders", bytes.NewBuffer(jsonValue))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.CreateOrder(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "service error")
		mockService.AssertExpectations(t)
	})

	// Test GetOrder
	t.Run("GetOrder - Success", func(t *testing.T) {
		mockService := new(MockOrderService)
		expectedOrder := &models.Order{ID: "order1", CustomerInfo: models.CustomerInfo{Name: "John Doe"}}
		mockService.On("GetOrder", "order1").Return(expectedOrder, nil)

		controller := &controllers.OrderController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "order_id", Value: "order1"}}

		controller.GetOrder(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "order1")
		mockService.AssertExpectations(t)
	})

	t.Run("GetOrder - Not Found", func(t *testing.T) {
		mockService := new(MockOrderService)
		mockService.On("GetOrder", "order1").Return(nil, errors.New("not found"))

		controller := &controllers.OrderController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "order_id", Value: "order1"}}

		controller.GetOrder(c)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, rr.Body.String(), "Order not found")
		mockService.AssertExpectations(t)
	})
}
