package models

import "gorm.io/gorm"

type Category struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	User    User
	UserID  string
	Deleted gorm.DeletedAt
}

func (u *Category) String() string {
	return u.Name
}
