package main

import (
	"Calc/agent"
	"Calc/db"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func main() {
	database := db.CreateDB()
	defer db.CreateDB().Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			login := r.Form.Get("login")
			password := r.Form.Get("password")
			if ConfirmNew(database, login) {
				SaveToDatabase(database, login, password)
				http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправление на страницу входа
				return
			} else {
				errorHTML := `<!DOCTYPE html>
<html>
<head>
    <title>Ошибка</title>
</head>
<body>
    <h1>Ошибка:</h1>
    <p>Такой пользователь уже существует!</p>
</body>
</html>`
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(errorHTML))
			}
		}
		html, err := os.ReadFile("html/registr.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			login := r.Form.Get("login")
			password := r.Form.Get("password")
			if !CompareDatas(database, login, password) {
				errorHTML := `<!DOCTYPE html>
<html>
<head>
    <title>Ошибка входа</title>
</head>
<body>
    <h1>Ошибка:</h1>
    <p>Пароль или Логин не совпадают!</p>
</body>
</html>`
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(errorHTML))
			} else {
				http.Redirect(w, r, "/calculator", http.StatusSeeOther)
			}
		}
		html, err := os.ReadFile("html/login.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	})

	http.HandleFunc("/calculator", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			expression := r.Form.Get("expression")
			time := r.Form.Get("time")
			timeInSeconds, _ := strconv.Atoi(time)
			// Преобразуем время из строки в число

			// Вызываем функцию ProcessExpression с указанным временем
			result := agent.ProcessExpression(expression, timeInSeconds)

			// Читаем файл с шаблоном результата
			htmlResult, err := os.ReadFile("html/result.html")
			if err != nil {
				http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError)
				return
			}

			// Устанавливаем заголовок Content-Type
			w.Header().Set("Content-Type", "text/html")

			// Создаем шаблон и передаем результат
			templateResult := template.Must(template.New("result").Parse(string(htmlResult)))
			templateResult.Execute(w, map[string]string{"Result": result})
			return
		}

		// Если запрос не POST, отдаем HTML форму для ввода выражения
		html, err := os.ReadFile("html/calc.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	})

	http.ListenAndServe(":8000", nil)
}

func SaveToDatabase(database *sql.DB, login, password string) error {
	ctx := context.Background()
	_, err := database.ExecContext(ctx, "INSERT INTO users (login, password) VALUES (?, ?)", login, password)
	return err
}

func CompareDatas(database *sql.DB, login, password string) bool {
	ctx := context.Background()
	row := database.QueryRowContext(ctx, "SELECT password FROM users WHERE login = ?", login)

	var dbPassword string
	err := row.Scan(&dbPassword)
	if err != nil {
		return false // Пользователь не найден
	}
	if password == dbPassword {
		return true // Пароли совпадают
	}

	return false // Пароли не совпадают
}

func ConfirmNew(database *sql.DB, login string) bool {
	ctx := context.Background()
	row := database.QueryRowContext(ctx, "SELECT login FROM users WHERE login = ?", login)
	var dbLogin string
	err := row.Scan(&dbLogin)
	if err != nil {
		return true // Пользователя с таким логином нет в базе
	}
	return false // Пользователь с таким логином уже существует
}
