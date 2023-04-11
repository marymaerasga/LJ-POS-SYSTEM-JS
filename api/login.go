package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/LJ-POS-SYSTEM-JS/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	dsn := "root:a@tcp(127.0.0.1:3306)/pos_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := models.User{}
	employee := []models.Employee{}

	db.Where("username = ?", username).Find(&user)
	db.Where("user_id = ?", user.ID).Find(&employee)

	if CheckPasswordHash(password, user.Password) {
		result := "1"

		newSession := uuid.NewString()

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "session",
			Value: newSession,
		})

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "id",
			Value: fmt.Sprint(user.ID),
		})

		http.SetCookie(w, &http.Cookie{
			Path:  "/",
			Name:  "try",
			Value: user.Position,
		})
		data := map[string]interface{}{
			"status":  "ok",
			"results": result,
		}
		ReturnJSON(w, r, data)
	} else {
		result := "0"
		data := map[string]interface{}{
			"status":  "error",
			"results": result,
		}
		ReturnJSON(w, r, data)
	}

}

func EditUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	user := models.User{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	username := r.FormValue("username")
	password := r.FormValue("password")
	

	db.Where("id", id).Find(&user)

	user.Username = username
	user.Password = hashPassword(password)

	db.Save(&user)

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	w.Write(b)
}
