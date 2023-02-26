package views

import (
	"net/http"
)

func StockHistoryHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Stock History | LJ POS_SYSTEM"

	return data

}
