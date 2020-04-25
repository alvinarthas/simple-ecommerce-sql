package config

import (
	"github.com/alvinarthas/simple-ecommerce-sql/models"
	"github.com/jinzhu/gorm"

	// set mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is for initialization Connection
var DB *gorm.DB

// InitDB is
func InitDB() {
	var err error

	// Setting Database MYSQL
	DB, err = gorm.Open("mysql", "root:@/simple_ecommerce?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")

	}

	// Migrate the Database

	// Customer
	DB.AutoMigrate(&models.Customer{})
	// Category
	DB.AutoMigrate(&models.Category{})
	// Store
	DB.AutoMigrate(&models.Store{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")

	// Related
	DB.Model(&models.Customer{}).Related(&models.Store{})

}
