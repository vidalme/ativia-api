package controllers

import (
	"fmt"
	"net/http"
)

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ola ativia")
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
	fmt.Println("Update um membro pelo nome")
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleta um membro pelo nome")
}
