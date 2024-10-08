package models

import (
	"github.com/vidalme/ativia-api/internal/db"
)

type User struct {
	Id         int
	UserName   string
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Phone      string
	UserStatus int
	// Id         int    `json:"id"`
	// UserName   string `json:"UserName"`
	// FirstName  string `json:"firstname"`
	// LastName   string `json:"lastname"`
	// Email      string `json:"email"`
	// Password   string `json:"password"`
	// Phone      string `json:"phone"`
	// UserStatus int    `json:"userstatus"`
}

func AdicionaUser() error {

	// var user User

	// user.Id = 123
	// user.UserName = "Fulano_DiTal"
	// user.FirstName = "Fulano"
	// user.LastName = "DiTal"
	// user.Email = "fulanodital@ativia.com"
	// user.Password = "qwe"
	// user.Phone = "23344455566"
	// user.UserStatus = 1

	// selectTodosUsers, err := db.Query("INSERT ")

	// 	INSERT INTO users (id, column2, column3, ...)
	// VALUES (value1, value2, value3, ...);

	db := db.ConectaComBancoDeDados()
	defer db.Close()

	return nil
}

func BuscaTodosUsers() []User {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectTodosUsers, err := db.Query("select * from users order by id asc")
	if err != nil {
		panic(err.Error())
	}

	u := User{}
	tu := []User{}

	for selectTodosUsers.Next() {
		var id, userStatus int
		var userName, firstName, lastName, email, password, phone string

		err := selectTodosUsers.Scan(&id, &userName, &firstName, &lastName, &email, &password, &phone, &userStatus)
		if err != nil {
			panic(err.Error())
		}
		u.Id = id
		u.UserStatus = userStatus
		u.UserName = userName
		u.FirstName = firstName
		u.LastName = lastName
		u.Email = email
		u.Password = password
		u.Phone = phone

		tu = append(tu, u)
	}

	return tu
}
