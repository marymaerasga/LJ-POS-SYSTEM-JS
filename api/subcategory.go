package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
)

func CreateSubCategory(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	category := models.SubCategory{}
	name := r.FormValue("name")
	user, _ := r.Cookie("id")
	id := r.FormValue("id")

	category.Name = name
	category.UserID = user.Value
	category.CategoryID = id
	db.Save(&category)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetSubCategory(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.SubCategory{}
	db.Preload("User").Preload("Category").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


func EditSubCategory(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	category := models.SubCategory{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	user, _ := r.Cookie("id")
	sub := r.FormValue("cat")

	db.Where("id", id).Find(&category)

	category.Name = name
	category.UserID = user.Value
	category.CategoryID = sub

	db.Save(&category)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.SubCategory{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

