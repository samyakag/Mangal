package tests

import (
	"errors"
	"testing"

	"mangal-chai-backend/models"
	"mangal-chai-backend/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(order models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrder(id string) (*models.Order, error) {
	args := m.Called(id)
	val := args.Get(0)
	if val == nil {
		return nil, args.Error(1)
	}
	return val.(*models.Order), args.Error(1)
}

type MockProductRepositoryForOrderService struct {
	mock.Mock
}

func (m *MockProductRepositoryForOrderService) GetProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepositoryForOrderService) GetProduct(id string) (*models.Product, error) {
	args := m.Called(id)
	val := args.Get(0)
	if val == nil {
		return nil, args.Error(1)
	}
	return val.(*models.Product), args.Error(1)
}

func (m *MockProductRepositoryForOrderService) GetProductsByCategory(category string) ([]models.Product, error) {
	args := m.Called(category)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepositoryForOrderService) GetCategories() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockProductRepositoryForOrderService) SeedProducts(products []interface{}) error {
	args := m.Called(products)
	return args.Error(0)
}

func TestOrderService(t *testing.T) {
	// Test CreateOrder
	t.Run("CreateOrder - Success", func(t *testing.T) {
		mockOrderRepo := new(MockOrderRepository)
		mockProductRepo := new(MockProductRepositoryForOrderService)

		product := &models.Product{ID: "prod1", Name: "Test Product", Price: 10.0, InStock: true}
		mockProductRepo.On("GetProduct", "prod1").Return(product, nil)
		mockOrderRepo.On("CreateOrder", mock.AnythingOfType("models.Order")).Return(nil)

		service := &services.OrderService{OrderRepository: mockOrderRepo, ProductRepository: mockProductRepo}

		orderData := struct {
			CustomerInfo models.CustomerInfo `json:"customer_info"`
			Items        []models.CartItem   `json:"items"`
			Notes        string              `json:"notes"`
		}{
			CustomerInfo: models.CustomerInfo{Name: "John Doe"},
			Items:        []models.CartItem{{ProductID: "prod1", Quantity: 1}},
			Notes:        "",
		}

		order, err := service.CreateOrder(orderData)

		assert.Nil(t, err)
		assert.NotNil(t, order)
		mockOrderRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("CreateOrder - Product Not Found", func(t *testing.T) {
		mockOrderRepo := new(MockOrderRepository)
		mockProductRepo := new(MockProductRepositoryForOrderService)

		mockProductRepo.On("GetProduct", "prod1").Return(nil, errors.New("not found"))

		service := &services.OrderService{OrderRepository: mockOrderRepo, ProductRepository: mockProductRepo}

		orderData := struct {
			CustomerInfo models.CustomerInfo `json:"customer_info"`
			Items        []models.CartItem   `json:"items"`
			Notes        string              `json:"notes"`
		}{
			CustomerInfo: models.CustomerInfo{Name: "John Doe"},
			Items:        []models.CartItem{{ProductID: "prod1", Quantity: 1}},
			Notes:        "",
		}

		order, err := service.CreateOrder(orderData)

		assert.NotNil(t, err)
		assert.Nil(t, order)
		mockOrderRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("CreateOrder - Product Out of Stock", func(t *testing.T) {
		mockOrderRepo := new(MockOrderRepository)
		mockProductRepo := new(MockProductRepositoryForOrderService)

		product := &models.Product{ID: "prod1", Name: "Test Product", Price: 10.0, InStock: false}
		mockProductRepo.On("GetProduct", "prod1").Return(product, nil)

		service := &services.OrderService{OrderRepository: mockOrderRepo, ProductRepository: mockProductRepo}

		orderData := struct {
			CustomerInfo models.CustomerInfo `json:"customer_info"`
			Items        []models.CartItem   `json:"items"`
			Notes        string              `json:"notes"`
		}{
			CustomerInfo: models.CustomerInfo{Name: "John Doe"},
			Items:        []models.CartItem{{ProductID: "prod1", Quantity: 1}},
			Notes:        "",
		}

		order, err := service.CreateOrder(orderData)

		assert.NotNil(t, err)
		assert.Nil(t, order)
		mockOrderRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	// Test GetOrder
	t.Run("GetOrder - Success", func(t *testing.T) {
		mockOrderRepo := new(MockOrderRepository)
		mockProductRepo := new(MockProductRepositoryForOrderService)

		expectedOrder := &models.Order{ID: "order1", CustomerInfo: models.CustomerInfo{Name: "John Doe"}}
		mockOrderRepo.On("GetOrder", "order1").Return(expectedOrder, nil)

		service := &services.OrderService{OrderRepository: mockOrderRepo, ProductRepository: mockProductRepo}
		order, err := service.GetOrder("order1")

		assert.Nil(t, err)
		assert.Equal(t, expectedOrder, order)
		mockOrderRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("GetOrder - Error", func(t *testing.T) {
		mockOrderRepo := new(MockOrderRepository)
		mockProductRepo := new(MockProductRepositoryForOrderService)

		mockOrderRepo.On("GetOrder", "order1").Return(nil, errors.New("not found"))

		service := &services.OrderService{OrderRepository: mockOrderRepo, ProductRepository: mockProductRepo}
		order, err := service.GetOrder("order1")

		assert.NotNil(t, err)
		assert.Nil(t, order)
		mockOrderRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})
}