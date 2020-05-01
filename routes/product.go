package routes

import (
	"strconv"

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
	// Get Store ID from Authorization token
	storeID := uint(c.MustGet("jwt_store_id").(float64))

	// set Parameter POST
	price, _ := strconv.Atoi(c.PostForm("price"))
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	weight, _ := strconv.Atoi(c.PostForm("weight"))
	condition, _ := strconv.ParseBool(c.PostForm("condition"))
	// category, err := strconv.ParseUint(c.PostForm("category"), 10, 32)
	// Get Form
	item := models.Product{
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		Price:        price,
		Condition:    condition,
		InitialStock: stock,
		Weight:       weight,
		StoreID:      storeID,
		// CategoryID:   category,
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "successfuly register user, please check your email",
		"data":   item,
	})
}

// UpdateProduct is to update existing product -> Store
func UpdateProduct(c *gin.Context) {
}

// DeleteProduct is to delete existing product -> Store
func DeleteProduct(c *gin.Context) {
}
