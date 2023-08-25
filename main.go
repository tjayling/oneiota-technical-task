package main

import (
	"bitbucket.org/oneiota/platform-technical-task/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Set the router.
	router := gin.Default()

	// Setup route group for the API.
	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("/", getProducts)
			products.GET("/:productID", getProduct)
		}
	}

	// Start and run the server
	router.Run(":3000")
}

func getProducts(c *gin.Context) {

	products := product_service.GetProducts()

	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	// Get product from CSV based on a given id.

	productId := c.Param("productID")
	product := product_service.GetProduct(productId)

	if product == nil {
		message := gin.H{
			"message": fmt.Sprintf("Product with id %s not found.", productId),
		}
		c.JSON(http.StatusNotFound, message)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
