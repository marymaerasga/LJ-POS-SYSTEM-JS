package models

import "gorm.io/gorm"

type SubCategory struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	User    User
	UserID  string
	Category Category
	CategoryID string
	Deleted gorm.DeletedAt
}

