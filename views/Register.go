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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Path:   "/",
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	})

	http.SetCookie(w, &http.Cookie{
		Path:   "/",
		Name:   "id",
		Value:  "",
		MaxAge: -1,
	})

	http.SetCookie(w, &http.Cookie{
		Path:   "/",
		Name:   "position",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
