
package main

import (
	"log"
	"os"
	"strings"

	"mangal-chai-backend/controllers"
	"mangal-chai-backend/database"
	"mangal-chai-backend/repositories"
	"mangal-chai-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getAllowedOrigins() []string {
	// Get allowed origins from environment variable
	// Format: comma-separated list of origins
	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")

	if allowedOriginsEnv != "" {
		origins := strings.Split(allowedOriginsEnv, ",")
		// Trim whitespace from each origin
		for i, origin := range origins {
			origins[i] = strings.TrimSpace(origin)
		}
		return origins
	}

	// Default allowed origins for development
	return []string{
		"http://localhost:5173",      // Vite dev server
		"http://localhost:3000",      // Alternative dev port
		"http://127.0.0.1:5173",
		"http://127.0.0.1:3000",
	}
}

func main() {
	// Database connection
	db := database.Connect()
	defer database.Disconnect()

	// Repositories
	productRepository := &repositories.ProductRepository{Collection: db.Collection("products")}
	orderRepository := &repositories.OrderRepository{Collection: db.Collection("orders")}

	// Services
	productService := &services.ProductService{Repository: productRepository}
	orderService := &services.OrderService{OrderRepository: orderRepository, ProductRepository: productRepository}
	paymentService := services.NewPaymentService()

	// Seed database
	if err := productService.SeedProducts(); err != nil {
		log.Fatal(err)
	}

	// Controllers
	productController := &controllers.ProductController{Service: productService}
	orderController := &controllers.OrderController{Service: orderService}
	paymentController := &controllers.PaymentController{Service: paymentService}

	// Gin router
	router := gin.Default()

	// CORS middleware - configure allowed origins
	allowedOrigins := getAllowedOrigins()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// API Routes
	api := router.Group("/api")
	{
		api.GET("/products", productController.GetProducts)
		api.GET("/products/:product_id", productController.GetProduct)
		api.GET("/products/category/:category", productController.GetProductsByCategory)
		api.POST("/orders", orderController.CreateOrder)
		api.GET("/orders/:order_id", orderController.GetOrder)
		api.GET("/categories", productController.GetCategories)
		api.POST("/payments/create-order", paymentController.CreateRazorpayOrder)
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "healthy", "message": "Mangal Chai API is running"})
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	log.Printf("Starting server on :%s", port)
	log.Printf("CORS enabled for origins: %v", allowedOrigins)
	router.Run(":" + port)
}
