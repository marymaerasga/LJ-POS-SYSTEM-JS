package views

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DashHandler(w http.ResponseWriter, r *http.Request) {
	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	session := r.FormValue("session")
	if session != "" {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: session,
			Path:  "/",
		})
	}

	user := models.User{}
	users, _ := r.Cookie("id")
	db.Where("id = ?", users.Value).Find(&user)

	activeSession := GetActiveSession(r)
	if activeSession == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {

		if(user.Position == "Staff"){

			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dash/")
			page := strings.TrimSuffix(r.URL.Path, "/")
			context := map[string]interface{}{}

			switch page {
			case "product":
				context = ProductHandler(w, r)
			case "stock-in-entry":
				context = StockEntryHandler(w, r)
			case "category":
				context = CategoryHandler(w, r)
			case "stock-in-history":
				context = StockHistoryHandler(w, r)
			case "product-stock-status":
				context = ProductStockHandler(w, r)
			case "pull-out-history":
				context = PulloutHistoryHandler(w, r)
			case "vat-discount":
				context = VatDiscountHandler(w, r)
			case "customer":
				context = CustomerHandler(w, r)
			case "admin-account":
				context = AdminAccountHandler(w, r)
			case "pos":
				context = POSHandler(w, r)
			default:
				page = "product"
			}

			ParseMultiHTML(w, r, page, context)

		} else if(user.Position == "Cashier"){
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dash/")
			page := strings.TrimSuffix(r.URL.Path, "/")
			context := map[string]interface{}{}

			switch page {
			case "admin-account":
				context = AdminAccountHandler(w, r)
			case "pos":
				context = POSHandler(w, r)
			default:
				page = "pos"
			}

			ParseMultiHTML(w, r, page, context)

		}else{

			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dash/")
			page := strings.TrimSuffix(r.URL.Path, "/")
			context := map[string]interface{}{}
	
			switch page {
			case "dashboard":
				context = DashboardHandler(w, r)
			case "product":
				context = ProductHandler(w, r)
			case "stock-in-entry":
				context = StockEntryHandler(w, r)
			case "category":
				context = CategoryHandler(w, r)
			case "stock-in-history":
				context = StockHistoryHandler(w, r)
			case "product-stock-status":
				context = ProductStockHandler(w, r)
			case "pull-out-history":
				context = PulloutHistoryHandler(w, r)
			case "vat-discount":
				context = VatDiscountHandler(w, r)
			case "customer":
				context = CustomerHandler(w, r)
			case "staff-account":
				context = StaffAccountHandler(w, r)
			case "admin-account":
				context = AdminAccountHandler(w, r)
			case "pos":
				context = POSHandler(w, r)
			default:
				page = "dashboard"
			}
	
			ParseMultiHTML(w, r, page, context)
		}
	
		}

	

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func ParseMultiHTML(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/base.html")

	path := "./templates/" + page + ".html"
	templateList = append(templateList, path)

	tmpl := template.Must(template.ParseFiles(templateList...))
	err := tmpl.ExecuteTemplate(w, "base.html", context)
	if err != nil {
		panic(err)
	}

}

func GetActiveSession(r *http.Request) *http.Cookie {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key
	}
	return nil
}
