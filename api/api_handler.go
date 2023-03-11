package api

import (
	"net/http"
	"strings"
)

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "employee") {
		CreateEmployee(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "login") {
		Login(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "item") {
		CreateItem(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_item") {
		GetItem(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_item") {
		EditItem(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_item") {
		DeleteItem(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "customer") {
		CreateCustomer(w, r)
		return
	}
}
