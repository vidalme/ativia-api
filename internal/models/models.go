package models

import (
	"github.com/vidalme/ativia-api/internal/db"
)

type User struct {
	Id         int
	Username   string
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Phone      string
	UserStatus int
	// Id         int    `json:"id"`
	// Username   string `json:"username"`
	// FirstName  string `json:"firstname"`
	// LastName   string `json:"lastname"`
	// Email      string `json:"email"`
	// Password   string `json:"password"`
	// Phone      string `json:"phone"`
	// UserStatus int    `json:"userstatus"`
}

func SelecionaTodosUsers() []User {
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
		u.Username = userName
		u.FirstName = firstName
		u.LastName = lastName
		u.Email = email
		u.Password = password
		u.Phone = phone

		tu = append(tu, u)
	}

	return tu
}
