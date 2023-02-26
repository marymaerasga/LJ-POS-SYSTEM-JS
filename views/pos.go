package views

import (
	"net/http"
)

func POSHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "POS | LJ POS_SYSTEM"

	return data

}
