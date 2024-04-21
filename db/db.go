package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() *sql.DB {
	ctx := context.TODO()
	database, err := sql.Open("sqlite3", "db/users.db")
	if err != nil {
		log.Fatal(err)
	}

	err = database.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS users (login TEXT, password TEXT expression TEXT, result INTEGER)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.ExecContext(ctx, "INSERT INTO users (login, password, expression, result) VALUES (?, ?, ?)")

	_, err = database.Query("SELECT login, password expression, result FROM users")

	if err != nil {
		log.Fatal(err)
	}
	return database
}
