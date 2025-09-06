package controllers

import (
	"mangal-chai-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Service services.ProductServiceInterface
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	products, err := c.Service.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	product, err := c.Service.GetProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) GetProductsByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	products, err := c.Service.GetProductsByCategory(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetCategories(ctx *gin.Context) {
	categories, err := c.Service.GetCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching categories"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}