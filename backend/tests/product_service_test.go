package tests

import (
	"errors"
	"testing"

	"mangal-chai-backend/models"
	"mangal-chai-backend/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) GetProduct(id string) (*models.Product, error) {
	args := m.Called(id)
	val := args.Get(0)
	if val == nil {
		return nil, args.Error(1)
	}
	return val.(*models.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductsByCategory(category string) ([]models.Product, error) {
	args := m.Called(category)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) GetCategories() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockProductRepository) SeedProducts(products []interface{}) error {
	args := m.Called(products)
	return args.Error(0)
}

func TestProductService(t *testing.T) {
	// Test GetProducts
	t.Run("GetProducts - Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		expectedProducts := []models.Product{{ID: "1", Name: "Test Product"}}
		mockRepo.On("GetProducts").Return(expectedProducts, nil)

		service := &services.ProductService{Repository: mockRepo}
		products, err := service.GetProducts()

		assert.Nil(t, err)
		assert.Equal(t, expectedProducts, products)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetProducts - Error", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("GetProducts").Return([]models.Product{}, errors.New("db error"))

		service := &services.ProductService{Repository: mockRepo}
		products, err := service.GetProducts()

		assert.NotNil(t, err)
		assert.Empty(t, products)
		mockRepo.AssertExpectations(t)
	})

	// Test GetProduct
	t.Run("GetProduct - Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		expectedProduct := &models.Product{ID: "1", Name: "Test Product"}
		mockRepo.On("GetProduct", "1").Return(expectedProduct, nil)

		service := &services.ProductService{Repository: mockRepo}
		product, err := service.GetProduct("1")

		assert.Nil(t, err)
		assert.Equal(t, expectedProduct, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetProduct - Error", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("GetProduct", "1").Return(nil, errors.New("not found"))

		service := &services.ProductService{Repository: mockRepo}
		product, err := service.GetProduct("1")

		assert.NotNil(t, err)
		assert.Nil(t, product)
		mockRepo.AssertExpectations(t)
	})

	// Test GetProductsByCategory
	t.Run("GetProductsByCategory - Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		expectedProducts := []models.Product{{ID: "1", Name: "Test Product", Category: "Tea"}}
		mockRepo.On("GetProductsByCategory", "Tea").Return(expectedProducts, nil)

		service := &services.ProductService{Repository: mockRepo}
		products, err := service.GetProductsByCategory("Tea")

		assert.Nil(t, err)
		assert.Equal(t, expectedProducts, products)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetProductsByCategory - Error", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("GetProductsByCategory", "Tea").Return([]models.Product{}, errors.New("db error"))

		service := &services.ProductService{Repository: mockRepo}
		products, err := service.GetProductsByCategory("Tea")

		assert.NotNil(t, err)
		assert.Empty(t, products)
		mockRepo.AssertExpectations(t)
	})

	// Test GetCategories
	t.Run("GetCategories - Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		expectedCategories := []string{"Tea", "Coffee"}
		mockRepo.On("GetCategories").Return(expectedCategories, nil)

		service := &services.ProductService{Repository: mockRepo}
		categories, err := service.GetCategories()

		assert.Nil(t, err)
		assert.Equal(t, expectedCategories, categories)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetCategories - Error", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("GetCategories").Return([]string{}, errors.New("db error"))

		service := &services.ProductService{Repository: mockRepo}
		categories, err := service.GetCategories()

		assert.NotNil(t, err)
		assert.Empty(t, categories)
		mockRepo.AssertExpectations(t)
	})

	// Test SeedProducts
	t.Run("SeedProducts - Success", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("SeedProducts", mock.Anything).Return(nil)

		service := &services.ProductService{Repository: mockRepo}
		err := service.SeedProducts()

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("SeedProducts - Error", func(t *testing.T) {
		mockRepo := new(MockProductRepository)
		mockRepo.On("SeedProducts", mock.Anything).Return(errors.New("db error"))

		service := &services.ProductService{Repository: mockRepo}
		err := service.SeedProducts()

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}