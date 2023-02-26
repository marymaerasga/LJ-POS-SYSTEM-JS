package views

import (
	"net/http"
)

func StaffAccountHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Staff Account | LJ POS_SYSTEM"

	return data

}
