package api

import (
	"fmt"
	"net/http"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}
	user := models.User{}
	employee := models.Employee{}

	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	username := r.FormValue("username")
	password := r.FormValue("password")

	user.FirstName = fname
	user.LastName = lname
	user.Username = username
	user.Password = hashPassword(password)
	db.Save(&user)

	employee.UserID = user.ID
	employee.EmployeeType = models.EmployeeType(2)
	db.Save(&employee)

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}
