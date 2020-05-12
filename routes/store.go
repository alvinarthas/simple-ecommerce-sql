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

// RegisterStore is to register
func RegisterStore(c *gin.Context) {

	// Get User ID from Authorization token
	userID := uint(c.MustGet("jwt_user_id").(float64))
	token, _ := RandomToken()

	// Get Form
	item := models.Store{
		Name:              c.PostForm("name"),
		UserName:          c.PostForm("user_name"),
		Adress:            c.PostForm("adress"),
		Email:             c.PostForm("email"),
		Phone:             c.PostForm("phone"),
		VerificationToken: token,
		UserID:            userID,
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err})
		c.Abort()
		return
	}

	var user models.User
	config.DB.Model(&user).Where("id = ?", userID).Updates(models.User{
		HaveStore: true,
	})

	// I want to send email for activation

	c.JSON(200, gin.H{
		"status": "successfuly register user, please check your email",
		"data":   item,
	})
}

// GetStore to gett all products that the store have
func GetStore(c *gin.Context) {
	// Get Store ID from Authorization token
	storeUsername := c.Param("username")

	// Set Query Params for Filtering & Sorting
	queryCategory := c.Query("category")
	querySort := c.Query("sort")
	queryPriceMin := c.Query("pricemin")
	queryPriceMax := c.Query("pricemax")
	queryCondition := c.Query("condition")

	var itemStore models.Store

	if config.DB.First(&itemStore, "user_name = ? AND is_activate = ?", storeUsername, 1).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	items := []models.Product{}

	query := config.DB.Where("store_id = ?", itemStore.ID)

	// Filter by Category
	if queryCategory != "" {
		query = query.Where("category_id = ?", queryCategory)
	}

	// Filter by Condition
	if queryCondition != "" {
		query = query.Where("`condition` = ?", queryCondition)
	}

	// Filter by Price
	if queryPriceMin != "" && queryPriceMax != "" {
		query = query.Where("price BETWEEN ? AND ?", queryPriceMin, queryPriceMax)
	} else if queryPriceMax != "" {
		query = query.Where("price <= ?", queryPriceMax)
	} else if queryPriceMin != "" {
		query = query.Where("price >= ?", queryPriceMin)
	}

	// Sorting
	if querySort != "" {
		if querySort == "high" {
			query = query.Order("price desc")
		} else if querySort == "low" {
			query = query.Order("price desc")
		} else if querySort == "atoz" {
			query = query.Order("name asc")
		} else if querySort == "ztoa" {
			query = query.Order("name desc")
		} else if querySort == "new" {
			query = query.Order("id desc")
		} else if querySort == "old" {
			query = query.Order("id asc")
		}
	}

	// Errors Tracing
	errors := query.Find(&items).GetErrors()

	if len(errors) > 0 {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": errors})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   items,
	})
}

// InfoStore to get store account info
func InfoStore(c *gin.Context) {
	// Get Store ID from Authorization token
	storeUsername := c.Param("username")

	var itemStore models.Store

	if config.DB.First(&itemStore, "user_name = ? AND is_activate = ?", storeUsername, 1).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   itemStore,
	})
}
