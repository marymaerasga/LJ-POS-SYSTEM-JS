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
	total := r.FormValue("total")
	user, _ := r.Cookie("id")
	date := r.FormValue("date")

	stock.ItemID = id
	stock.Stock = total
	stock.UserID = user.Value
	stock.Date = date
	db.Save(&stock)


	item := models.Item{}
	db.Where("id", id).Find(&item)

	//  temp,_ := strconv.Atoi(r.FormValue("total"))

	// item.Quantity += uint(temp)
	db.Save(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetStockin(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.StockIn{}
	db.Preload("Item").Find(&item)

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
	item := models.Item{}
	id, _ := strconv.Atoi(r.FormValue("itemID"))
	sid, _ := strconv.Atoi(r.FormValue("stockID"))
	// current, _ := strconv.Atoi(r.FormValue("CurrentCount"))
	// new, _ := strconv.Atoi(r.FormValue("NewCount"))

	db.Where("id", id).Find(&item)
	// item.Quantity -= uint(current)
	// item.Quantity += uint(new)
	db.Save(&item)

	stockin := models.StockIn{}
	latest := r.FormValue("NewCount")
	db.Where("id", sid).Find(&stockin)
	stockin.Stock = latest;
	db.Save(&stockin)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteStockin(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	item := models.Item{}
	id, _ := strconv.Atoi(r.FormValue("itemID"))
	sid, _ := strconv.Atoi(r.FormValue("stockID"))
	// count, _ := strconv.Atoi(r.FormValue("new"))

	db.Where("id", id).Find(&item)
	// item.Quantity -= uint(count)
	db.Save(&item)


	stockin := models.StockIn{}
	db.Where("id", sid).Statement.Delete(&stockin)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
