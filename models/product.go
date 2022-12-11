package models

type Item struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       uint
	Quantity    uint
	Size        Size
	Category    Category
	Image       string
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

type Category int

func (Category) Category1() int {
	return 1
}
func (Category) Category2() int {
	return 2
}
func (Category) Category3() int {
	return 3
}

func (e *Item) Save() {
}
