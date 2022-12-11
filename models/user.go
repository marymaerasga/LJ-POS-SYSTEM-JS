package models

type User struct {
	ID        uint
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func (u *User) String() string {
	return u.FirstName + " " + u.LastName
}
