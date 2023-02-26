package views

import (
	"net/http"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Product | LJ POS_SYSTEM"

	return data

}
