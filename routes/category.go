package routes

import (
	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/models"
	"github.com/gin-gonic/gin"
)

/*
	Category for CRUD , and to get all
*/

// GetAllCategories is to get all category -> Admin Only
func GetAllCategories(c *gin.Context) {
	items := []models.Category{}
	config.DB.Find(&items)

	// Return JSON
	c.JSON(200, gin.H{
		"status":  "berhasil",
		"message": "Berhasil menampilkan semua data category",
		"data":    items,
	})
}

// GetCategory is to get spesific product -> Store
func GetCategory(c *gin.Context) {
	// Get Parameter
	id := c.Param("id")

	var item models.Category

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

// CreateCategory is to create new category
func CreateCategory(c *gin.Context) {

	// Get Form
	item := models.Category{
		Name: c.PostForm("name"),
		Desc: c.PostForm("desc"),
		Icon: c.PostForm("icon"),
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "Successfully Create Category",
		"data":   item,
	})
}

// UpdateCategory is to update existing product -> Store
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var item models.Category

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	config.DB.Model(&item).Where("id = ?", id).Updates(models.Category{
		Name: c.PostForm("name"),
		Desc: c.PostForm("desc"),
		Icon: c.PostForm("icon"),
	})

	c.JSON(200, gin.H{
		"status": "berhasil update data category",
		"data":   item,
	})
}

// DeleteCategory is to delete existing category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var item models.Category

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	config.DB.Where("id = ?", id).Delete(&item)
	c.JSON(200, gin.H{
		"status": "berhasil delete",
		"data":   item,
	})
}
