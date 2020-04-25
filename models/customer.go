package models

import "github.com/jinzhu/gorm"

// Customer is model for the customer
type Customer struct {
	gorm.Model
	Stores   []Store // to show that customer can have many stores
	UserName string
	FullName string
	Email    string
	SocialID string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
}
