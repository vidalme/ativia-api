package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "host=postgres port=5432 user=admin password=qwe dbname=postgres sslmode=disable "
	// connStr := "user=postgres password=yourpassword dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error)
	}
	return db
}
