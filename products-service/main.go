package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/***** Models *****/

// product represents data about a record product.
type product struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

// products slice to seed record product data.
var products = []product{
	{ID: "1", Title: "Pizza 1", Price: 56.99},
	{ID: "2", Title: "Pizza 2", Price: 17.99},
	{ID: "3", Title: "Pizza 3", Price: 39.99},
}

/***** Router *****/

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", postProducts)

	router.Run("localhost:8080")
}

/***** API Handlers *****/

// getProducts responds with the list of all products as JSON.
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// postProducts adds a product from JSON received in the request body.
func postProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new product to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

// getProductByID locates the product whose ID value matches the id
// parameter sent by the client, then returns that product as a response.
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of products, looking for
	// a product whose ID value matches the parameter.
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
