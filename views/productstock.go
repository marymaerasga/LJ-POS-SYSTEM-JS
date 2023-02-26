package views

import (
	"net/http"
)

func ProductStockHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Product Stock | LJ POS_SYSTEM"

	return data

}
