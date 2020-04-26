package main

import (
	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/routes"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	// set up database
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	//  Setting Default Router
	router := gin.Default()

	// Initialize Version
	apiV1 := router.Group("/api/v1/")
	{
		// Normal Register and Login
		apiV1.POST("/register", routes.RegisterUser)
		apiV1.POST("/login", routes.LoginUser)

		// Users
		users := apiV1.Group("/users")
		{
			// Initilize Http method for Users Crud
			users.GET("/", routes.GetUser)
			users.GET("/:id", routes.GetUserByID)
		}
	}

	router.Run()
}
