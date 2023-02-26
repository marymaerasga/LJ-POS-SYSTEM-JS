package views

import (
	"net/http"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Category | LJ POS_SYSTEM"

	return data

}
