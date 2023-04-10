package models

import "gorm.io/gorm"

type Customer struct {
	ID       uint `gorm:"primaryKey"`
	User     User
	UserID   string
	Name     string
	MobileNo string
	Address  string
	Deleted  gorm.DeletedAt
}
