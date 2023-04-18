package api

import (
	"encoding/json"
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
	category := r.FormValue("category")
	sub := r.FormValue("sub")
	user, _ := r.Cookie("id")
	expiration := r.FormValue("expiration")
	alert := r.FormValue("alert")
	low := r.FormValue("low")
	

	item.Name = name
	item.Description = description
	item.CategoryID = category
	item.SubCategoryID = sub
	item.ExpiredDate = alert
	item.Low = low
	item.Expiration = expiration
	item.UserID = user.Value
	item.Status = "0"
	item.Code = "I" + category + name[0:2] ;

	db.Save(&item)
	db.Where("id = ?", item.ID).Find(&item)

	

	new_order := r.FormValue("cart")
	var c []map[string]string
	json.Unmarshal([]byte(new_order), &c)

	for i := range c {
		product := models.ProductItem{}
		fmt.Println(c[i]["size"])
		data, _ := strconv.Atoi(c[i]["size"])

		

		product.ItemID = int(item.ID)
		product.Size = models.Size(data)
		product.Color = c[i]["color"]
		product.PurchasePrice = c[i]["p_price"]
		product.RetailedPrice = c[i]["r_price"]
		product.Quantity = 0
		db.Save(&product)

	}

}

func GetSubItem(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	

	item := []models.ProductItem{}
	db.Preload("Item").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

}
