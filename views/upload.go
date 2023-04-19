package views

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/upload.html"))
	data := map[string]interface{}{}
	data["Title"] = "Upload Image | LJ POS_SYSTEM"

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	if r.Method == "POST" {
		item := models.Item{}
		id,_ := strconv.Atoi(r.FormValue("itemID"))
		db.Where("id", id).Find(&item)

		mydir, _ := os.Getwd()
		b := mydir + "/static/img/item/"
		file, handler, err := r.FormFile("productImages")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)

		}

		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		dst, err := os.Create(b + handler.Filename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}


		file1, handler1, err := r.FormFile("productImages1")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)

		}

		defer file1.Close()
		fmt.Printf("Uploaded File: %+v\n", handler1.Filename)
		fmt.Printf("File Size: %+v\n", handler1.Size)
		fmt.Printf("MIME Header: %+v\n", handler1.Header)

		dst1, err := os.Create(b + handler1.Filename)
		defer dst1.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		if _, err := io.Copy(dst1, file1); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}


		file2, handler2, err := r.FormFile("productImages2")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)

		}

		defer file2.Close()
		fmt.Printf("Uploaded File: %+v\n", handler2.Filename)
		fmt.Printf("File Size: %+v\n", handler2.Size)
		fmt.Printf("MIME Header: %+v\n", handler2.Header)

		dst2, err := os.Create(b + handler2.Filename)
		defer dst2.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		if _, err := io.Copy(dst2, file2); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		item.Image = "/static/img/item/" + handler.Filename
		item.Image2 = "/static/img/item/" + handler1.Filename
		item.Image3 = "/static/img/item/" + handler2.Filename

		db.Save(&item)

		http.Redirect(w, r, "/dash/product", http.StatusSeeOther)
	}
	tmpl.Execute(w, data)
	
}
