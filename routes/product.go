package routes

import (
	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/models"
	"github.com/gin-gonic/gin"
)

/*
	User can have more than one store
	User still can shop and User still can sell their things in their stores
*/

// GetAllProducts is to get all prodcuts -> Admin Only
func GetAllProducts(c *gin.Context) {
}

// GetProduct is to get spesific product -> Store
func GetProduct(c *gin.Context) {
	// Get Parameter
	id := c.Param("id")

	var item models.Product

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}

// CreateProduct is to create new product -> Store
func CreateProduct(c *gin.Context) {
}

// UpdateProduct is to update existing product -> Store
func UpdateProduct(c *gin.Context) {
}

// DeleteProduct is to delete existing product -> Store
func DeleteProduct(c *gin.Context) {
}
