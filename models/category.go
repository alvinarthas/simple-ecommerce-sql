package models

import "github.com/jinzhu/gorm"

// Category to map the product
type Category struct {
	gorm.Model
	Name string
	Desc string
	Icon string
}
