package models

import (
	"github.com/vidalme/ativia-api/internal/db"
)

type User struct {
	Id         int    `json:"id"`
	UserName   string `json:"username"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userstatus"`
}

func GetAllUsers() []User {
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

func AddUser(userName, firstName, lastName, email, password, phone string, userStatus int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	userDataSTMT, err := db.Prepare("insert into users(username, firstname, lastname, email, password, phone, userstatus) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}
	userDataSTMT.Exec(userName, firstName, lastName, email, password, phone, userStatus)
}

func AddUsers(users []User) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	userDataSTMT, err := db.Prepare("insert into users(username, firstname, lastname, email, password, phone, userstatus) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}

	for _, u := range users {
		userDataSTMT.Exec(u.UserName, u.FirstName, u.LastName, u.Email, u.Password, u.Phone, u.UserStatus)
	}

}

func DeleteUser(userName string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	userDataSTMT, err := db.Prepare("delete from users where username=$1")
	if err != nil {
		panic(err.Error())
	}
	userDataSTMT.Exec(userName)
}

func SelectUser(un string) User {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectUserQuery := "select * from users where username=$1"
	row := db.QueryRow(selectUserQuery, un)

	var id, userStatus int
	var userName, firstName, lastName, email, password, phone string

	err := row.Scan(&id, &userName, &firstName, &lastName, &email, &password, &phone, &userStatus)
	if err != nil {
		panic(err.Error())
	}

	u := User{}
	u.Id = id
	u.UserName = userName
	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
	u.Password = password
	u.Phone = phone
	u.UserStatus = userStatus

	return u
}

func EditUser(oldUserName string, userStatusConvInt int, user User) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	editUserQuery := "update users set username=$1, firstname=$2, lastname=$3, email=$4, password=$5, phone=$6, userstatus=$7 where username=$8;"
	editUser, err := db.Prepare(editUserQuery)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("a variavel user.UserStatus é do tipo %T e tem o valor de %d\n", user.UserStatus, user.UserStatus)
	_, err = editUser.Exec(user.UserName, user.FirstName, user.LastName, user.Email, user.Password, user.Phone, userStatusConvInt, oldUserName)
	if err != nil {
		panic(err.Error())
	}

}
