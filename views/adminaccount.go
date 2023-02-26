package views

import (
	"net/http"
)

func AdminAccountHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Admin Account | LJ POS_SYSTEM"

	return data

}
