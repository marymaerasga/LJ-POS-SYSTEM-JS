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

	if strings.HasPrefix(r.URL.Path, "order") {
		CreateOrder(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_employee") {
		GetEmployee(w, r)
		return
	}

	
	if strings.HasPrefix(r.URL.Path, "edit_employee") {
		EditEmployee(w, r)
		return
	}

	
	if strings.HasPrefix(r.URL.Path, "delete_employee") {
		DeleteEmployee(w, r)
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

	if strings.HasPrefix(r.URL.Path, "expired") {
		ExpiredItem(w, r)
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

	if strings.HasPrefix(r.URL.Path, "get_customer") {
		GetCustomer(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_customer") {
		EditCustomer(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_customer") {
		DeleteCustomer(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "category") {
		CreateCategory(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_category") {
		GetCategory(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "edit_category") {
		EditCategory(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "delete_category") {
		DeleteCategory(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "stockin") {
		CreateStockin(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_stockin") {
		GetStockin(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_stockin") {
		EditStockin(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_stockin") {
		DeleteStockin(w, r)
		return
	}
}
