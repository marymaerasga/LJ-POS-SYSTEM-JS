package api

import (
	"fmt"
	"net/http"

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
	var user_id = 1
	customer := models.Customer{}

	name := r.FormValue("name")
	mobileNo := r.FormValue("mobileNo")
	address := r.FormValue("address")


	customer.Name = name
	customer.MobileNo = mobileNo
	customer.Address = address
	customer.UserID = user_id
	db.Save(&customer)
}


