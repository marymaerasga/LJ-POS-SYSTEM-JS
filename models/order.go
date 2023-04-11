package models

import "gorm.io/gorm"

type Order struct {
	ID         uint `gorm:"primaryKey"`
	Date       string
	Name		string
	Discount   float64
	Vat        float64
	SubTotal   float64
	GrandTotal float64
	User       User
	UserID     string
	Payment    string
	Change		string
	Deleted    gorm.DeletedAt
}

func (u *Order) uint() uint {
	return u.ID
}
