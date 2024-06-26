package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
)

func CreateStockin(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	stock := models.StockIn{}
	id := r.FormValue("id")
	total,_ := strconv.Atoi(r.FormValue("total"))
	user, _ := r.Cookie("id")
	date := r.FormValue("date")

	stock.ProductItemID = id
	stock.Stock = total
	stock.UserID = user.Value
	stock.Date = date
	db.Save(&stock)


	item := models.ProductItem{}
	db.Where("id", id).Find(&item)
	println("result",item.Quantity)

	 temp,_ := strconv.Atoi(r.FormValue("total"))

	item.Quantity += int(temp)
	db.Save(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetStockin(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.StockIn{}
	db.Preload("ProductItem").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func EditStockin(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	item := models.ProductItem{}
	id, _ := strconv.Atoi(r.FormValue("itemID"))
	sid, _ := strconv.Atoi(r.FormValue("stockID"))
	current, _ := strconv.Atoi(r.FormValue("CurrentCount"))
	new, _ := strconv.Atoi(r.FormValue("NewCount"))



	db.Where("id", id).Find(&item)
	item.Quantity -= int(new)
	item.Quantity += int(current)
	db.Save(&item)

	stockin := models.StockIn{}
	latest,_ := strconv.Atoi(r.FormValue("CurrentCount"))
	db.Where("id", sid).Find(&stockin)
	stockin.Stock = latest;
	db.Save(&stockin)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteStockin(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	item := models.ProductItem{}
	id, _ := strconv.Atoi(r.FormValue("itemID"))
	sid, _ := strconv.Atoi(r.FormValue("stockID"))
	count, _ := strconv.Atoi(r.FormValue("new"))

	db.Where("id", id).Find(&item)
	item.Quantity -= int(count)
	db.Save(&item)


	stockin := models.StockIn{}
	db.Where("id", sid).Statement.Delete(&stockin)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
