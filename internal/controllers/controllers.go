package controllers

import (
	"fmt"
	"net/http"

	"github.com/vidalme/ativia-api/internal/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ola ativia 0.0.6 mais uma vez")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cria User")
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
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Deleta um membro pelo nome")
}

// so para testes
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := models.SelecionaTodosUsers()
	fmt.Fprintln(w, "todos os users:", allUsers)
}
