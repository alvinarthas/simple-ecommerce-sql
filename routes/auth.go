package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// RegisterUser to store the new customer data into DB
func RegisterUser(c *gin.Context) {
	// Check Password confirmation
	password := c.PostForm("password")
	confirmedPassword := c.PostForm("confirmed_password")

	// Return Error if not confirmed
	if password != confirmedPassword {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "password not confirmed"})
		c.Abort()
		return
	}

	// Hash the password
	hash, _ := HashPassword(password)

	// Get Form
	item := models.User{
		UserName: c.PostForm("user_name"),
		FullName: c.PostForm("full_name"),
		Email:    c.PostForm("email"),
		Password: hash,
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err})
		c.Abort()
		return
	}

	// I want to send email for activation

	c.JSON(200, gin.H{
		"status": "successfuly register user, please check your email",
		"data":   item,
	})
}

// LoginUser to get the token for access the system
func LoginUser(c *gin.Context) {
	// Get Login Form
	email := c.PostForm("email")
	password := c.PostForm("password")

	item := models.User{}

	// Check if email and password is match
	if !config.DB.First(&item, "email = ?", email).RecordNotFound() && CheckPasswordHash(password, item.Password) {
		token := createToken(&item)

		c.JSON(200, gin.H{
			"status": "success",
			"data":   item,
			"token":  token,
		})

	} else {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "your email/password may be wrong",
		})
	}
}

// CreateToken to generate token for accesing the system
func createToken(user *models.User) string {
	// to send time expire, issue at (iat)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"user_role":  user.Role,
		"user_store": user.HaveStore,
		"exp":        time.Now().AddDate(0, 0, 7).Unix(),
		"iat":        time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
