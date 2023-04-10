package api

import (
	"fmt"
	"net/http"
	"strconv"

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
	position := r.FormValue("position")
	sex := r.FormValue("sex")
	number := r.FormValue("mobileNo")
	address := r.FormValue("address")


	user.FirstName = fname
	user.LastName = lname
	user.Username = username
	user.Password = hashPassword(password)
	user.Position = position
	db.Save(&user)

	employee.UserID = user.ID
	employee.Address = address
	employee.Sex = sex
	employee.MobileNo = number
	employee.Status = "Active"
	db.Save(&employee)

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Employee{}
	db.Preload("User").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func EditEmployee(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	user := models.User{}
	employee := models.Employee{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	username := r.FormValue("username")
	status := r.FormValue("status")
	position := r.FormValue("position")
	sex := r.FormValue("sex")
	number := r.FormValue("mobileNo")
	address := r.FormValue("address")

	db.Where("username", username).Find(&user)
	user.FirstName = fname
	user.LastName = lname
	user.Username = username
	user.Position = position
	db.Save(&user)


	db.Where("id", id).Find(&employee)

	employee.Address = address
	employee.Sex = sex
	employee.MobileNo = number
	employee.Status = status
	db.Save(&employee)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	id, _ := strconv.Atoi(r.FormValue("id"))
	uid, _ := strconv.Atoi(r.FormValue("uid"))
	item := models.Employee{}
	db.Where("id", id).Statement.Delete(&item)

	user := models.User{}
	db.Where("id", uid).Statement.Delete(&user)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
