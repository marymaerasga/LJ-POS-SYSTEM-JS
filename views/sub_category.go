package views

import (
	"net/http"
)

func SubCategoryHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Category | LJ POS_SYSTEM"

	return data

}
