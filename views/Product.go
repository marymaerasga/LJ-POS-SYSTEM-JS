package views

import (
	"net/http"
	"text/template"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/product.html"))
	data := map[string]interface{}{}
	data["Title"] = "PRODUCT | LJ POS_SYSTEM"
	tmpl.Execute(w, data)
}
