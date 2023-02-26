package views

import (
	"net/http"
)

func StockEntryHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Stock Entry | LJ POS_SYSTEM"

	return data

}
