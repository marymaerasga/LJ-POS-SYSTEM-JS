package views

import (
	"net/http"
)

func PulloutHistoryHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	data["Title"] = "Pullout History | LJ POS_SYSTEM"

	return data

}
