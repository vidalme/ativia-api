package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	// usando temporariamente os ips absolutos por que estou compilando o codigo fora do docker compose,
	// entao fora da network sem possibilidade de usar o DNS interno do compose
	// conexao := "host=172.18.0.2 port=5432 user=admin password=qwe dbname=postgres sslmode=disable "
	conexao := "host=172.19.0.2 port=5432 user=admin password=qwe dbname=postgres sslmode=disable "
	// connStr := "user=postgres password=yourpassword dbname=yourdbname sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error)
	}
	return db
}
