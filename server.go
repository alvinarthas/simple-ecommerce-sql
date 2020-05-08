package main

import (
	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/middleware"
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
		// Social Auth or OAuth
		apiV1.GET("/auth/:provider", routes.RedirectHandler)
		apiV1.GET("/auth/:provider/callback", routes.CallbackHandler)

		// Normal Register and Login
		apiV1.POST("/register", routes.RegisterUser)
		apiV1.POST("/login", routes.LoginUser)

		// Verification User and Store Account
		apiV1.GET("/verify/store/:token", routes.VerifyStoreAccount)
		apiV1.GET("/verify/user/:token", routes.VerifyUserAccount)

		// User
		user := apiV1.Group("/user")
		{
			// Initilize Http method for User Crud
			user.GET("/", routes.GetUser)
			user.GET("/:id", routes.GetUserByID)
		}

		// Category
		category := apiV1.Group("/category")
		{
			// Initilize Http method for Category Crud
			category.GET("/", routes.GetAllCategories)
			category.GET("/:id", middleware.IsAdmin(), routes.GetCategory)
			category.POST("/create", middleware.IsAdmin(), routes.CreateCategory)
			category.PUT("/update/:id", middleware.IsAdmin(), routes.UpdateCategory)
			category.DELETE("/delete/:id", middleware.IsAdmin(), routes.DeleteCategory)
		}

		// Store
		store := apiV1.Group("/store")
		{
			store.POST("/register", middleware.IsAuth(), routes.RegisterStore)
			store.GET("/:username", routes.GetStore)                               //show all store products
			store.GET("/:username/info", middleware.HaveStore(), routes.InfoStore) // Account Info
		}

		// Product CRUD by Store
		product := apiV1.Group("/product")
		{
			product.GET("/", routes.GetAllProducts) // every product in every store
			product.GET("/:id", routes.GetProduct)
			product.POST("/create", middleware.HaveStore(), routes.CreateProduct)
			product.PUT("/update/:id", middleware.HaveStore(), routes.UpdateProduct)
			product.DELETE("/delete/:id", middleware.HaveStore(), routes.DeleteProduct)
		}
	}

	router.Run()
}
