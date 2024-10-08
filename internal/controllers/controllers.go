package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/vidalme/ativia-api/internal/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Ola ativia 0.0.6 mais uma vez")
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

func AddUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Adiciona multiplos novos membros")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userName := r.PathValue("userName")
	fmt.Fprintf(w, "Rertorna um membro pelo nome: %v\n", userName)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update um membro pelo nome")
}

// so para testes
// func GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	allUsers := models.BuscaTodosUsers()
// 	fmt.Fprintln(w, "todos os users:", allUsers)
// }
