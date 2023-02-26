package views

import (
	"net/http"
)

func CustomerHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Customer | LJ POS_SYSTEM"

	return data

}
