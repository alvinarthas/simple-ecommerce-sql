package models

import "github.com/jinzhu/gorm"

// User is model for the customer
type User struct {
	gorm.Model
	Stores     []Store // to show that customer can have many stores
	UserName   string
	FullName   string
	Email      string `gorm:"unique_index"`
	Password   string
	SocialID   string
	Provider   string
	Avatar     string
	Role       bool `gorm:"default:0"`
	HaveStore  bool `gorm:"default:0"`
	isActivate bool `gorm:"default:0"`
}
