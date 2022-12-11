package models

type Employee struct {
	ID           uint `gorm:"primaryKey"`
	User         User
	UserID       uint
	EmployeeType EmployeeType
}

type EmployeeType int

func (EmployeeType) Admin() int {
	return 1
}
func (EmployeeType) User() int {
	return 2
}

func (e *Employee) Save() {
}

func (e *Employee) String() string {
	return e.User.String()
}
