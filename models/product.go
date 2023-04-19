package models

import "gorm.io/gorm"

type Item struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	User    User
	UserID  string
	Category    Category
	CategoryID  string
	SubCategory SubCategory
	SubCategoryID string
	Expiration string
	ExpiredDate string
	Low			string
	Image       string
	Image2       string
	Image3       string
	Status		string
	Code 		string
	Deleted gorm.DeletedAt
}


