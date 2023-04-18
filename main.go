package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/api"
	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"github.com/dafalo/LJ-POS-SYSTEM-JS/views"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":2027"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	CreateDB("pos_system")
	MigrateDB()
	CreateDefaultUser()
	Handlers()

	http.ListenAndServe(Port, nil)
}

func Handlers() {
	fmt.Println("Handlers")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", views.LoginHandler)
	http.HandleFunc("/register", views.RegisterHandler)
	http.HandleFunc("/dash/", views.DashHandler)
	http.HandleFunc("/api/", api.APIHandler)
	http.HandleFunc("/logout", views.LogoutHandler)

}

func CreateDB(name string) *sql.DB {
	fmt.Println("Database Created")
	db, err := sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func MigrateDB() {
	fmt.Println("Database Migrated")
	employees := models.Employee{}
	user := models.User{}
	category := models.Category{}
	item := models.Item{}
	customer := models.Customer{}
	stockin := models.StockIn{}
	order := models.Order{}
	orderlines := models.Orderlines{}
	subcategory := models.SubCategory{}
	subitem := models.ProductItem{}

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&employees, &user, &item, &customer, &category,&stockin,&order,&orderlines,&subcategory,&subitem)
}

func CreateDefaultUser() {
	fmt.Println("User")
	db := GormDB()

	user := []models.User{}
	db.Find(&user)

	defaultUser := []models.User{
		{
			Username: "Master@yahoo.com",
			Password: hashPassword("Master"),
			FirstName: "Master",
			LastName: "Master",
			Position: "Admin",
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Username == users.Username {
				isExisting = true
				break
			}
		}

		if !isExisting {
			fmt.Println("Create Default User")
			db.Save(&defaultUser[i])
		}
	}

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func GormDB() *gorm.DB {
	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}