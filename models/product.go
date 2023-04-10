package models

import "gorm.io/gorm"

type Item struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       uint
	Quantity    uint
	Size        Size
	User    User
	UserID  string
	Category    Category
	CategoryID  string
	Expiration string
	Image       string
	Deleted gorm.DeletedAt
}

type Size int

func (Size) ExtraSmall() int {
	return 1
}
func (Size) Small() int {
	return 2
}
func (Size) Medium() int {
	return 3
}
func (Size) Large() int {
	return 4
}
func (Size) ExtraLarge() int {
	return 5
}


