package views

import (
	"net/http"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	data := map[string]interface{}{}
	data["Title"] = "Main | LJ POS_SYSTEM"
	tmpl.Execute(w, data)
}
