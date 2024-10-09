package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/vidalme/ativia-api/internal/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", models.GetAllUsers())
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	userName := r.FormValue("user_name")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	phone := r.FormValue("phone")

	models.AddUser(userName, firstName, lastName, email, password, phone, 1)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	userName := r.PathValue("userName")
	models.DeleteUser(userName)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userName := r.PathValue("userName")
	tmpl.ExecuteTemplate(w, "User", models.SelectUser(userName))
}

func AddUsers(w http.ResponseWriter, r *http.Request) {
	// gin transforms in a single method
	users := []models.User{}
	data, _ := io.ReadAll(r.Body)
	json.Unmarshal(data, &users)
	// ----

	models.AddUsers(users)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update um membro pelo nome")
}
