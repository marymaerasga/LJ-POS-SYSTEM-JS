package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetItem(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	

	item := []models.Item{}
	db.Preload("Category").Where("status", "0").Find(&item)

	expired := []models.Item{}
	db.Preload("Category").Where("status", "1").Find(&expired)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"expired": expired,
	}
	ReturnJSON(w, r, data)

}

func ExpiredItem(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	

	item := models.Item{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	db.Preload("Category").Where("id", id).Find(&item)

	item.Status = "1";

	db.Save(&item)

}

func EditItem(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	description := r.FormValue("description")
	qty, _ := strconv.Atoi(r.FormValue("quantity"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	size, _ := strconv.Atoi(r.FormValue("size"))
	category :=r.FormValue("category")

	item := models.Item{}
	db.Where("id", id).Find(&item)
	item.Name = name
	item.Description = description
	item.Quantity = uint(qty)
	item.Price = uint(price)
	item.Size = models.Size(size)
	item.CategoryID = category

	db.Save(&item)

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Item{}
	db.Where("id", id).Statement.Delete(&item)

}
