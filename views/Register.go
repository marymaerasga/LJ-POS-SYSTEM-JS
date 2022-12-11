package views

import (
	"net/http"
	"text/template"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/register.html"))
	data := map[string]interface{}{}
	data["Title"] = "Register | LJ POS_SYSTEM"
	tmpl.Execute(w, data)
}
