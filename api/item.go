package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	new_order := r.FormValue("cart")
	subtotal, _ := strconv.ParseFloat(r.FormValue("subtotal"), 64)
	discount, _ := strconv.ParseFloat(r.FormValue("discount"), 64)
	vat, _ := strconv.ParseFloat(r.FormValue("vat"), 64)
	total, _ := strconv.ParseFloat(r.FormValue("total"), 64)
	cash := r.FormValue("cash")
	change := r.FormValue("change")
	name := r.FormValue("name")
	current := time.Now()
	date := current.Format("2006-02-01")
	user, _ := r.Cookie("id")

	// * Initialize the models needed
	orders := models.Order{}

	// * Convert the JSON into an array of map
	var c []map[string]string
	json.Unmarshal([]byte(new_order), &c)

	// * Create order
	orders.Date = date
	orders.SubTotal = float64(subtotal)
	orders.Discount = float64(discount)
	orders.Vat = float64(vat)
	orders.GrandTotal = float64(total)
	orders.UserID = user.Value
	orders.Payment = cash
	orders.Change = change
	orders.Name = name

	db.Save(&orders)
	db.Where("id = ?", orders.ID).Find(&orders)

	// * Declare the struct for the passing of values
	for i := range c {
		orderlines := models.Orderlines{}
		oiID := 0
		oiQTY, _ := strconv.ParseFloat("0", 64)
		oiID, _ = strconv.Atoi(c[i]["itemid"])
		oiQTY, _ = strconv.ParseFloat(c[i]["quantity"], 64)

		println(oiQTY)

		orderlines.OrderID = orders.ID
		orderlines.Date = date
		orderlines.ItemID = uint(oiID)
		orderlines.Quantity = float64(oiQTY)
		orderlines.UserID = user.Value
		db.Save(&orderlines)

		product := models.Item{}
		db.Where("id", oiID).Find(&product)
		product.Quantity = uint(product.Quantity) - uint(oiQTY)
		db.Save(&product)
	}

	res := map[string]interface{}{
		"status": "ok",
	}
	ReturnJSON(w, r, res)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func GetOrder(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	item := []models.Order{}
	db.Preload("User").Find(&item)

	orderlines := []models.Orderlines{}
	db.Preload("Order").Preload("Item").Find(&orderlines)

	data := map[string]interface{}{
		"status": "ok",
		"order":   item,
		"orderlines": orderlines,
	}
	ReturnJSON(w, r, data)

}


func SearchOrder(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	item := []models.Order{}
	to := r.FormValue("to")
	from := r.FormValue("from")
	db.Preload("User").Where("date BETWEEN ? and ?", from, to).Find(&item)

	orderlines := []models.Orderlines{}
	db.Preload("Order").Preload("Item").Where("date BETWEEN ? and ?", from, to).Find(&orderlines)

	data := map[string]interface{}{
		"status": "ok",
		"order":   item,
		"orderlines": orderlines,
	}
	ReturnJSON(w, r, data)

}
