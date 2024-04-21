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
 database := db.CreateDB() // Создание соединения с базой данных
 defer db.CreateDB().Close() // Закрытие соединения с базой данных при завершении программы

 // Обработчик для корневого пути ("/")
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost { // Проверка, является ли запрос POST-запросом
   r.ParseForm()                 // Парсинг формы из запроса
   login := r.Form.Get("login")    // Получение логина из формы
   password := r.Form.Get("password") // Получение пароля из формы

   if ConfirmNew(database, login) { // Проверка, существует ли уже пользователь с таким логином
    SaveToDatabase(database, login, password) // Сохранение нового пользователя в базу данных
    http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправление на страницу входа
    return
   } else {
    // Формирование HTML-страницы с ошибкой
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
    w.Header().Set("Content-Type", "text/html") // Установка типа содержимого ответа
    w.Write([]byte(errorHTML)) // Запись HTML-кода в ответ
   }
  }
  // Чтение файла с HTML-шаблоном для страницы регистрации
  html, err := os.ReadFile("html/registr.html")
  if err != nil {
   http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError) // Обработка ошибки чтения файла
   return
  }
  w.Header().Set("Content-Type", "text/html") // Установка типа содержимого ответа
  w.Write(html)                                // Запись html-кода в ответ
 })

 // Обработчик для пути "/login"
 http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
  if r.Method == http.MethodPost { // Проверка, является ли запрос POST-запросом
   r.ParseForm()                 // Парсинг формы из запроса
   login := r.Form.Get("login")    // Получение логина из формы
   password := r.Form.Get("password") // Получение пароля из формы

   if !CompareDatas(database, login, password) { // Проверка соответствия логина и пароля
    // Формирование HTML-страницы с ошибкой
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
    w.Header().Set("Content-Type", "text/html") // Установка типа содержимого ответа 
    w.Write([]byte(errorHTML)) // Запись HTML-кода в ответ
   } else { 
    http.Redirect(w, r, "/calculator", http.StatusSeeOther) // Перенаправление на страницу калькулятора
   }
  } 
  // Чтение файла с HTML-шаблоном для страницы входа 
  html, err := os.ReadFile("html/login.html") 
  if err != nil {
   http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError) // Обработка ошибки чтения файла
   return
  }
  w.Header().Set("Content-Type", "text/html") // Установка типа содержимого ответа
  w.Write(html)                                // Запись html-кода в ответ 
 })

// Обработчик для пути "/calculator"
http.HandleFunc("/calculator", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost { // Проверка, является ли запрос POST-запросом
      r.ParseForm() // Парсинг формы из запроса
      expression := r.Form.Get("expression") // Получение выражения из формы
      time := r.Form.Get("time")            // Получение времени из формы
      timeInSeconds, _ := strconv.Atoi(time) // Преобразование времени из строки в число

      // Вызываем функцию ProcessExpression с указанным временем
      result := agent.ProcessExpression(expression, timeInSeconds)

      // Читаем файл с шаблоном результата
      htmlResult, err := os.ReadFile("html/result.html")
      if err != nil {
        http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError) // Обработка ошибки чтения файла
        return
      }

      // Устанавливаем заголовок Content-Type
      w.Header().Set("Content-Type", "text/html")

      // Создаем шаблон и передаем результат
      templateResult := template.Must(template.New("result").Parse(string(htmlResult)))
      templateResult.Execute(w, map[string]string{"Result": result})
      return
    }

    // Если запрос не POST, отдаем html форму для ввода выражения
    html, err := os.ReadFile("html/calc.html")
    if err != nil {
      http.Error(w, fmt.Sprintf("Ошибка чтения файла: %v", err), http.StatusInternalServerError) // Обработка ошибки чтения файла
      return
    }
    w.Header().Set("Content-Type", "text/html") // Установка типа содержимого ответа
    w.Write(html)                                // Запись html-кода в ответ
  })

// Запуск HTTP-сервера на порту 8000
http.ListenAndServe(":8000", nil)
}

// Функция для сохранения данных пользователя в базу данных
func SaveToDatabase(database *sql.DB, login, password string) error {
  ctx := context.Background()
  _, err := database.ExecContext(ctx, "INSERT INTO users (login, password) VALUES (?, ?)", login, password) // Выполнение SQL-запроса на вставку
  return err // Возвращение ошибки, если она возникла
}

// Функция для сравнения данных пользователя с данными из базы данных
func CompareDatas(database *sql.DB, login, password string) bool {
  ctx := context.Background()
  row := database.QueryRowContext(ctx, "SELECT password FROM users WHERE login = ?", login) // Выполнение SQL-запроса на выборку

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

// Функция для проверки, является ли пользователь новым 
func ConfirmNew(database *sql.DB, login string) bool {
  ctx := context.Background()
  row := database.QueryRowContext(ctx, "SELECT login FROM users WHERE login = ?", login) // Выполнение SQL-запроса на выборку
  var dbLogin string
  err := row.Scan(&dbLogin)
  if err != nil {
    return true // Пользователя с таким логином нет в базе
  }
  return false // Пользователь с таким логином уже существует
}
