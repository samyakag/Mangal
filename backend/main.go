
package main

import (
	"log"

	"mangal-chai-backend/controllers"
	"mangal-chai-backend/database"
	"mangal-chai-backend/repositories"
	"mangal-chai-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	// Seed database
	if err := productService.SeedProducts(); err != nil {
		log.Fatal(err)
	}

	// Controllers
	productController := &controllers.ProductController{Service: productService}
	orderController := &controllers.OrderController{Service: orderService}

	// Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
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
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "healthy", "message": "Mangal Chai API is running"})
		})
	}

	log.Println("Starting server on :8080")
	router.Run(":8001")
}
