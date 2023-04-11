package models

import "gorm.io/gorm"

type Orderlines struct {
	ID        uint `gorm:"primaryKey"`
	Date      string
	Order     Order
	OrderID   uint
	Item      Item
	ItemID	 uint
	Quantity  float64
	User      User
	UserID    string
	Deleted   gorm.DeletedAt
}

func (u *Orderlines) uint() uint {
	return u.ID
}
