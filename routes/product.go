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
	items := []models.Product{}
	config.DB.Find(&items)

	// Return JSON
	c.JSON(200, gin.H{
		"status":  "berhasil",
		"message": "Berhasil menampilkan semua data product",
		"data":    items,
	})
}

// GetStoreProducts to gett all products that the store have
func GetStoreProducts(c *gin.Context) {
	// Get Store ID from Authorization token
	storeID := uint(c.MustGet("jwt_store_id").(float64))

	items := []models.Product{}

	if config.DB.Find(&items, "store_id = ?", storeID).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   items,
	})
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
	categoryRaw, _ := strconv.ParseUint(c.PostForm("category"), 10, 32)

	// change category to uint
	category := uint(categoryRaw)

	// Get Form
	item := models.Product{
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		Price:        price,
		Condition:    condition,
		InitialStock: stock,
		Weight:       weight,
		StoreID:      storeID,
		CategoryID:   category,
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
	id := c.Param("id")

	var item models.Product

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	storeID := uint(c.MustGet("jwt_store_id").(float64))
	if storeID != item.StoreID {
		c.JSON(403, gin.H{
			"status":  "error",
			"message": "this data is forbidden"})
		c.Abort()
		return
	}

	// Get Form for updating the product
	price, _ := strconv.Atoi(c.PostForm("price"))
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	weight, _ := strconv.Atoi(c.PostForm("weight"))
	condition, _ := strconv.ParseBool(c.PostForm("condition"))
	categoryRaw, _ := strconv.ParseUint(c.PostForm("category"), 10, 32)

	// change category to uint
	category := uint(categoryRaw)

	config.DB.Model(&item).Where("id = ?", id).Updates(models.Product{
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		Price:        price,
		Condition:    condition,
		InitialStock: stock,
		Weight:       weight,
		StoreID:      storeID,
		CategoryID:   category,
	})

	c.JSON(200, gin.H{
		"status": "berhasil update data product",
		"data":   item,
	})
}

// DeleteProduct is to delete existing product -> Store
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var item models.Product

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	storeID := uint(c.MustGet("jwt_store_id").(float64))
	if storeID != item.StoreID {
		c.JSON(403, gin.H{
			"status":  "error",
			"message": "this data is forbidden"})
		c.Abort()
		return
	}

	config.DB.Where("id = ?", id).Delete(&item)
	c.JSON(200, gin.H{
		"status": "berhasil delete",
		"data":   item,
	})
}
