package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}
	item := models.Item{}

	name := r.FormValue("name")
	description := r.FormValue("description")
	qty, _ := strconv.Atoi(r.FormValue("quantity"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	size, _ := strconv.Atoi(r.FormValue("size"))
	category := r.FormValue("category")
	user, _ := r.Cookie("id")
	expiration := r.FormValue("expiration")
	

	item.Name = name
	item.Description = description
	item.Quantity = uint(qty)
	item.Price = uint(price)
	item.Size = models.Size(size)
	item.CategoryID = category
	item.Expiration = expiration
	item.UserID = user.Value
	item.Status = "0"
	item.Code = "I" + category + name[0:2] ;

	db.Save(&item)

}
