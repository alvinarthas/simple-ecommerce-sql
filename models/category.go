package models

import "github.com/jinzhu/gorm"

// Category to map the product
type Category struct {
	gorm.Model
	Products []Product // to show that Category can have many products
	Name     string
	Desc     string
	Icon     string
	Slug     string `gorm:"unique_index"`
}
