package models

type Customer struct {
	ID           uint `gorm:"primaryKey"`
	User         User
	UserID       int
	Name         string
	MobileNo     string
	Address      string
}

