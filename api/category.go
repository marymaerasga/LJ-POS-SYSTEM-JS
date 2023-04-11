package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	category := models.Category{}
	name := r.FormValue("name")
	user, _ := r.Cookie("id")

	category.Name = name
	category.UserID = user.Value
	db.Save(&category)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetCategory(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Category{}
	db.Preload("User").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.User{}
	user, _ := r.Cookie("id")
	db.Where("id", user.Value).Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func EditCategory(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	category := models.Category{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	user, _ := r.Cookie("id")

	db.Where("id", id).Find(&category)

	category.Name = name
	category.UserID = user.Value

	db.Save(&category)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Category{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GormDB() *gorm.DB {
	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}