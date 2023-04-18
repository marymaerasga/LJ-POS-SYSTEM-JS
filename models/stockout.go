package models

import "gorm.io/gorm"

type StockOut struct {
	ID          uint `gorm:"primaryKey"`
	User    	User
	UserID  	string
	StockIn    	StockIn
	StockInID  	int
	QTY			string
	Remarks		string
	Date       string
	Deleted gorm.DeletedAt
}


