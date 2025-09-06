package tests

import (
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

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetProducts() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) GetProduct(id string) (*models.Product, error) {
	args := m.Called(id)
	val := args.Get(0)
	if val == nil {
		return nil, args.Error(1)
	}
	return val.(*models.Product), args.Error(1)
}

func (m *MockProductService) GetProductsByCategory(category string) ([]models.Product, error) {
	args := m.Called(category)
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) GetCategories() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockProductService) SeedProducts() error {
	args := m.Called()
	return args.Error(0)
}

func TestProductController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Test GetProducts
	t.Run("GetProducts - Success", func(t *testing.T) {
		mockService := new(MockProductService)
		expectedProducts := []models.Product{{ID: "1", Name: "Test Product"}}
		mockService.On("GetProducts").Return(expectedProducts, nil)

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		controller.GetProducts(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Test Product")
		mockService.AssertExpectations(t)
	})

	t.Run("GetProducts - Error", func(t *testing.T) {
		mockService := new(MockProductService)
		mockService.On("GetProducts").Return([]models.Product{}, errors.New("service error"))

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		controller.GetProducts(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "Error fetching products")
		mockService.AssertExpectations(t)
	})

	// Test GetProduct
	t.Run("GetProduct - Success", func(t *testing.T) {
		mockService := new(MockProductService)
		expectedProduct := &models.Product{ID: "1", Name: "Test Product"}
		mockService.On("GetProduct", "1").Return(expectedProduct, nil)

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "product_id", Value: "1"}}

		controller.GetProduct(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Test Product")
		mockService.AssertExpectations(t)
	})

	t.Run("GetProduct - Not Found", func(t *testing.T) {
		mockService := new(MockProductService)
		mockService.On("GetProduct", "1").Return(nil, errors.New("not found"))

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "product_id", Value: "1"}}

		controller.GetProduct(c)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, rr.Body.String(), "Product not found")
		mockService.AssertExpectations(t)
	})

	// Test GetProductsByCategory
	t.Run("GetProductsByCategory - Success", func(t *testing.T) {
		mockService := new(MockProductService)
		expectedProducts := []models.Product{{ID: "1", Name: "Test Product", Category: "Tea"}}
		mockService.On("GetProductsByCategory", "Tea").Return(expectedProducts, nil)

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "category", Value: "Tea"}}

		controller.GetProductsByCategory(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Test Product")
		mockService.AssertExpectations(t)
	})

	t.Run("GetProductsByCategory - Error", func(t *testing.T) {
		mockService := new(MockProductService)
		mockService.On("GetProductsByCategory", "Tea").Return([]models.Product{}, errors.New("service error"))

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "category", Value: "Tea"}}

		controller.GetProductsByCategory(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "Error fetching products")
		mockService.AssertExpectations(t)
	})

	// Test GetCategories
	t.Run("GetCategories - Success", func(t *testing.T) {
		mockService := new(MockProductService)
		expectedCategories := []string{"Tea", "Coffee"}
		mockService.On("GetCategories").Return(expectedCategories, nil)

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		controller.GetCategories(c)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "Tea")
		mockService.AssertExpectations(t)
	})

	t.Run("GetCategories - Error", func(t *testing.T) {
		mockService := new(MockProductService)
		mockService.On("GetCategories").Return([]string{}, errors.New("service error"))

		controller := &controllers.ProductController{Service: mockService}

		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		controller.GetCategories(c)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "Error fetching categories")
		mockService.AssertExpectations(t)
	})
}