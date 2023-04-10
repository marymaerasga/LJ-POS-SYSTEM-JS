package models

import "gorm.io/gorm"

type Employee struct {
	ID           uint `gorm:"primaryKey"`
	User         User
	UserID       uint
	Sex			string
	MobileNo	string
	Address		string
	Status		string
	Deleted gorm.DeletedAt

}

func (e *Employee) String() string {
	return e.User.String()
}
