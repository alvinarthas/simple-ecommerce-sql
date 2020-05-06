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

	// User
	DB.AutoMigrate(&models.User{})
	// Category
	DB.AutoMigrate(&models.Category{})
	// Store
	DB.AutoMigrate(&models.Store{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// Product
	DB.AutoMigrate(&models.Product{}).AddForeignKey("store_id", "stores(id)", "CASCADE", "CASCADE").AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")

	// Related
	DB.Model(&models.User{}).Related(&models.Store{})
	DB.Model(&models.Store{}).Related(&models.Product{})
	DB.Model(&models.Category{}).Related(&models.Product{})

}
