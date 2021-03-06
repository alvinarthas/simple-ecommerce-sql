package routes

import (
	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/models"
	"github.com/gin-gonic/gin"
)

// GetUser to view all customers
func GetUser(c *gin.Context) {
	items := []models.User{}
	if config.DB.Find(&items).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "sucess",
		"data":   items,
	})
}

// GetUserByID to get the user data based on the user id
func GetUserByID(c *gin.Context) {
	// Get Parameter
	id := c.Param("id")

	var item models.User

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
