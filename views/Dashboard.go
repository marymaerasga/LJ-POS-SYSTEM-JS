package views

import (
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Dashboard | LJ POS_SYSTEM"

	return data

}

