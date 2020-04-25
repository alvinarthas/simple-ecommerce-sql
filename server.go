package main

import (
	"github.com/alvinarthas/simple-ecommerce-sql/config"
	"github.com/alvinarthas/simple-ecommerce-sql/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// set up database
	config.InitDB()
	defer config.DB.Close()

	//  Setting Default Router
	router := gin.Default()

	// Initialize Version
	apiV1 := router.Group("/api/v1/")
	{
		// Customers
		customers := apiV1.Group("/customers")
		{
			// Initilize Http method for Customers Crud
			customers.GET("/", routes.GetCustomer)
			customers.GET("/:id", routes.GetCustomerByID)
			customers.POST("/store", routes.PostCustomer)
			// customers.PUT("/update/:id", middleware.IsAuth(), routes.UpdateCustomer)
			// customers.DELETE("/delete/:id", middleware.IsAuth(), routes.DeleteCustomer)
		}
	}

	router.Run()
}
