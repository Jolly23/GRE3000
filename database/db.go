package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	PgSslMode     = "disable"
	PgHostAddress = "link-docker-postgresql"
	PgHostPort    = "5432"
	DatabaseName  = "GRE3000"
	PgUserName    = "JOLLY"
	PgPassword    = "200900"
)

type Database struct {
	conn *sql.DB
}

func NewDatabase() *Database {
	url := fmt.Sprintf("sslmode=%s host=%s port=%s dbname=%s user=%s password=%s",
		PgSslMode,
		PgHostAddress,
		PgHostPort,
		DatabaseName,
		PgUserName,
		PgPassword,
	)

	conn, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(); err != nil {
		panic(err)
	}
	return &Database{conn: conn}
}
