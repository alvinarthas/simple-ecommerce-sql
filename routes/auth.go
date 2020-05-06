package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/models"
	"github.com/danilopolani/gocialite/structs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// RegisterUser to store the new customer data into DB
func RegisterUser(c *gin.Context) {
	// Check Password confirmation
	password := c.PostForm("password")
	confirmedPassword := c.PostForm("confirmed_password")
	token, _ := RandomToken()

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
		UserName:          c.PostForm("user_name"),
		FullName:          c.PostForm("full_name"),
		Email:             c.PostForm("email"),
		Password:          hash,
		VerificationToken: token,
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
	if !config.DB.First(&item, "email = ? AND is_activate = ?", email, 1).RecordNotFound() && CheckPasswordHash(password, item.Password) {
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

// RedirectHandler to correct oAuth URL
func RedirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("CLIENT_ID_GH"),
			"clientSecret": os.Getenv("CLIENT_SECRET_GH"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/github/callback",
		},
		"google": {
			"clientID":     os.Getenv("CLIENT_ID_G"),
			"clientSecret": os.Getenv("CLIENT_SECRET_G"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/google/callback",
		},
	}

	providerScopes := map[string][]string{
		"github": []string{"public_repo"},
		"google": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// CallbackHandler Handle callback of provider
func CallbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	// Handle callback and check for errors
	user, _, err := config.Gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	var newUser = getOrRegisterUser(provider, user)
	var jtwToken = createToken(&newUser)

	c.JSON(200, gin.H{
		"data":    newUser,
		"token":   jtwToken,
		"message": "berhasil login",
	})
}

// Register new social ID to database
func getOrRegisterUser(provider string, user *structs.User) models.User {
	var userData models.User

	config.DB.Where("provider = ? AND social_id = ?", provider, user.ID).First(&userData)

	if userData.ID == 0 {
		token, _ := RandomToken()

		newUser := models.User{
			FullName:          user.FullName,
			UserName:          user.Username,
			Email:             user.Email,
			SocialID:          user.ID,
			Provider:          provider,
			Avatar:            user.Avatar,
			VerificationToken: token,
		}

		config.DB.Create(&newUser)

		return newUser
	}

	return userData
}

// CreateToken to generate token for accesing the system
func createToken(user *models.User) string {
	var store models.Store
	var storeID uint

	if user.HaveStore == true {
		if config.DB.First(&store, "user_id = ?", user.ID).RecordNotFound() {
			storeID = 0
		}
		storeID = store.ID
	} else {
		storeID = 0
	}
	// to send time expire, issue at (iat)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"user_role":  user.Role,
		"user_store": user.HaveStore,
		"store_id":   storeID,
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

// VerifyUserAccount to verify user and store account
func VerifyUserAccount(c *gin.Context) {

	verificationToken := c.Param("token")

	var item models.User

	if config.DB.First(&item, "verification_token = ?", verificationToken).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	config.DB.Model(&item).Where("id = ?", item.ID).Updates(models.User{
		IsActivate: true,
	})

	c.JSON(200, gin.H{
		"status": "Success, your account is now active. Please Login to your account",
		"data":   item,
	})
}

// VerifyStoreAccount to verify user and store account
func VerifyStoreAccount(c *gin.Context) {

	verificationToken := c.Param("token")

	var itemUser models.User
	var itemStore models.Store

	if config.DB.First(&itemStore, "verification_token = ?", verificationToken).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found"})
		c.Abort()
		return
	}

	if config.DB.First(&itemUser, "id = ?", itemStore.UserID).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "user record not found"})
		c.Abort()
		return
	}

	config.DB.Model(&itemUser).Where("id = ?", itemUser.ID).Updates(models.User{
		IsActivate: true,
		HaveStore:  true,
	})

	config.DB.Model(&itemStore).Where("id = ?", itemStore.ID).Updates(models.Store{
		IsActivate: true,
	})

	c.JSON(200, gin.H{
		"status": "Success, your account is now active. Please Login to your account",
	})

}
