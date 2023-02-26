package views

import (
	"net/http"
)

func VatDiscountHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Vat Discount | LJ POS_SYSTEM"

	return data

}

