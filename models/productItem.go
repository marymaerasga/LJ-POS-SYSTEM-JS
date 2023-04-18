package models

import "gorm.io/gorm"

type ProductItem struct {
	ID          uint `gorm:"primaryKey"`
	Item 	Item
	ItemID 	int
	Color 	string
	Size        Size
	PurchasePrice string
	RetailedPrice string

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
func (Size) None() int {
	return 6
}



