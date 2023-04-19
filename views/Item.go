package views

import (
	"net/http"
	"text/template"
)

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/receipt.html"))
	data := map[string]interface{}{}
	data["Title"] = "Receipt | LJ POS_SYSTEM"
	tmpl.Execute(w, data)
	
}
