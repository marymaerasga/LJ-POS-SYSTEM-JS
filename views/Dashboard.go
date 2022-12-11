package views

import (
	"net/http"
	"text/template"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/dashboard.html"))
	data := map[string]interface{}{}
	data["Title"] = "Dashboard | LJ POS_SYSTEM"
	tmpl.Execute(w, data)
}
