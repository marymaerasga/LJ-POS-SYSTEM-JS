package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
)

func CreateStocOut(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	stock := models.StockOut{}
	current1 := time.Now()
	date := current1.Format("2006-02-01")
	id, _ := strconv.Atoi(r.FormValue("itemID"))
	sid, _ := strconv.Atoi(r.FormValue("stockID"))
	current := r.FormValue("qty")
	remarks := r.FormValue("remarks")
	user, _ := r.Cookie("id")
	

	stock.StockInID = int(sid)
	stock.QTY = current
	stock.UserID = user.Value
	stock.Remarks = remarks
	stock.Date = date
	db.Save(&stock)


	item := models.ProductItem{}
	db.Where("id", id).Find(&item)
	 temp,_ := strconv.Atoi(r.FormValue("qty"))

	item.Quantity -= int(temp)
	db.Save(&item)

	stockin := models.StockIn{}
	db.Where("id", sid).Find(&stockin)
	 temp1,_ := strconv.Atoi(r.FormValue("qty"))

	stockin.Stock -= int(temp1)
	db.Save(&stockin)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetStockOut(w http.ResponseWriter, r *http.Request) {

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

func EditStockOut(w http.ResponseWriter, r *http.Request) {

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
	// latest := r.FormValue("CurrentCount")
	db.Where("id", sid).Find(&stockin)
	// stockin.Stock = latest;
	db.Save(&stockin)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func DeleteStockOut(w http.ResponseWriter, r *http.Request) {
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
