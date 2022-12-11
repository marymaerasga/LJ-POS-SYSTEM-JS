package api

import (
	"fmt"
	"net/http"

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
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

}
