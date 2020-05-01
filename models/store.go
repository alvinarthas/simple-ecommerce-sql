package models

import "github.com/jinzhu/gorm"

/*
Store is for the sellers,
User can have more than one stores
*/

// Store model
type Store struct {
	gorm.Model
	Products          []Product // to show that store can have many products
	Name              string
	UserName          string
	Adress            string
	Email             string
	Phone             string
	Avatar            string
	IsActivate        bool `gorm:"default:0"`
	VerificationToken string
	UserID            uint
}
