package models

import "gorm.io/gorm"

type StockIn struct {
	ID          uint `gorm:"primaryKey"`
	User    	User
	UserID  	string
	ProductItem    	ProductItem
	ProductItemID  	string
	Stock 		string
	Date       string
	Deleted gorm.DeletedAt
}


