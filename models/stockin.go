package models

import "gorm.io/gorm"

type StockIn struct {
	ID          uint `gorm:"primaryKey"`
	User    	User
	UserID  	string
	Item    	Item
	ItemID  	string
	Stock 		string
	Date       string
	Deleted gorm.DeletedAt
}


