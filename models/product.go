package models

import "github.com/jinzhu/gorm"

/*
Product is Belong To Store
One Store can have many products
One Product can only have one Store
One Product can only have one Categories
*/

// Product model
type Product struct {
	gorm.Model
	Name         string
	Description  string
	Condition    string
	Price        int
	InitialStock int
	Weight       int
	StoreID      uint
	CategoryID   uint
}
