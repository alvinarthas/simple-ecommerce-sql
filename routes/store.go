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

	// Get Form
	item := models.Store{
		Name:     c.PostForm("name"),
		UserName: c.PostForm("user_name"),
		Adress:   c.PostForm("address"),
		Email:    c.PostForm("email"),
		Phone:    c.PostForm("phone"),
		UserID:   userID,
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err})
		c.Abort()
		return
	}
}
