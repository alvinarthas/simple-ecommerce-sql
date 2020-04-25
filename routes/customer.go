package routes

import (
	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/models"
	"github.com/gin-gonic/gin"
)

// GetCustomer to view all customers
func GetCustomer(c *gin.Context) {
	items := []models.Customer{}
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

// PostCustomer to store the new customer data into DB
func PostCustomer(c *gin.Context) {
	// Get Form
	item := models.Customer{
		UserName: c.PostForm("user_name"),
		FullName: c.PostForm("full_name"),
		Email:    c.PostForm("email"),
	}

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil store data customer",
		"data":   item,
	})
}

// GetCustomerByID to get the customer data based on the username
func GetCustomerByID(c *gin.Context) {
	// Get Parameter
	id := c.Param("id")

	item := []models.Customer{}

	if config.DB.First(&item, "id = ?", id).RecordNotFound() || len(item) == 0 {
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
