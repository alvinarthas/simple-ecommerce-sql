package models

import "github.com/jinzhu/gorm"

/*
Store is for the sellers,
User can have more than one stores
*/

// Store model
type Store struct {
	gorm.Model
	Name   string
	Adress string
	Email  string
	Phone  string
	Avatar string
	UserID uint
}
