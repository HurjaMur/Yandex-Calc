package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Функция для создания и инициализации базы данных
func CreateDB() *sql.DB {
 ctx := context.TODO() // Создание контекста

 // Открытие базы данных sqlite3
 database, err := sql.Open("sqlite3", "db/users.db")
 if err != nil {
  log.Fatal(err) // Вывод фатальной ошибки, если не удалось открыть базу данных
 }

 // Проверка соединения с базой данных
 err = database.PingContext(ctx)
 if err != nil {
  log.Fatal(err) // Вывод фатальной ошибки, если не удалось подключиться к базе данных
 }

 // Создание таблицы "users", если она не существует
 _, err = database.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS users (login TEXT, password TEXT, expression TEXT, result INTEGER)")
 if err != nil {
  log.Fatal(err) // Вывод фатальной ошибки, если не удалось создать таблицу
 } 
	
  _, err = database.ExecContext(ctx, "INSERT INTO users (login, password, expression, result) VALUES (?, ?, ?)")
	
  _, err = database.Query("SELECT login, password expression, result FROM users")

  if err != nil {
    log.Fatal(err)
  }
 return database // Возвращение объекта базы данных
}
