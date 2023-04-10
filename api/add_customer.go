package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}
	user, _ := r.Cookie("id")
	customer := models.Customer{}

	name := r.FormValue("name")
	mobileNo := r.FormValue("mobileNo")
	address := r.FormValue("address")


	customer.Name = name
	customer.MobileNo = mobileNo
	customer.Address = address
	customer.UserID = user.Value
	db.Save(&customer)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Customer{}
	db.Preload("User").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func EditCustomer(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	category := models.Customer{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	number := r.FormValue("number")
	address := r.FormValue("address")

	db.Where("id", id).Find(&category)

	category.Name = name
	category.MobileNo= number
	category.Address = address

	db.Save(&category)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Customer{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


