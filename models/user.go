package models

import "github.com/jinzhu/gorm"

// User is model for the customer
type User struct {
	gorm.Model
	Store             Store // to show that customer can have one store
	UserName          string
	FullName          string
	Email             string `gorm:"unique_index"`
	Password          string `json:"-"`
	SocialID          string `json:"-"`
	Provider          string `json:"-"`
	Avatar            string
	Role              bool `gorm:"default:0"`
	HaveStore         bool `gorm:"default:0"`
	IsActivate        bool `gorm:"default:0"`
	VerificationToken string
}
