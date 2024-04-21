package agent

import (
	"strconv"
	"strings"
	"time"
)

// Функция для вычисления результата выражения и обновления статуса и значения
func ProcessExpression(expression string, tim int) string {
    t := time.Duration(tim) * time.Second // Преобразование времени в секунды

    resultChan := make(chan string, 1) // Создание канала для передачи результата

    // Запуск горутины для вычисления выражения
    go func() {
        time.Sleep(t)                       // Ожидание заданного времени
        resultChan <- calculate(expression) // Отправка результата в канал
    }()

    return <-resultChan // Возвращение результата из канала
}

// Функция для вычисления арифметического выражения
func calculate(expression string) string {
    tokens := strings.FieldsFunc(expression, func(r rune) bool {
        return r == '+' || r == '-' || r == '*' || r == '/' // Разделение выражения на токены
    })

    nums := make([]int, 0)         // Массив для хранения чисел
    operations := make([]string, 0) // Массив для хранения операций

    // Парсинг чисел из токенов
    for _, token := range tokens {
        num, err := strconv.Atoi(token)
        if err != nil {
            return "Ошибка: неверный формат числа" // Возвращение ошибки, если не удалось преобразовать токен в число
        }
        nums = append(nums, num)
    }

    // Парсинг операций из выражения
    for _, c := range expression {
        char := string(c)
        if char == "+" || char == "-" || char == "*" || char == "/" {
            operations = append(operations, char)
        }
    }

    // Выполнение умножения и деления
    for i := 0; i < len(operations); i++ {
        if operations[i] == "*" {
            nums[i] *= nums[i+1]
            nums = append(nums[:i+1], nums[i+2:]...)
            operations = append(operations[:i], operations[i+1:]...)
            i--
        } else if operations[i] == "/" {
            if nums[i+1] == 0 {
                return "Ошибка: деление на ноль" // Возвращение ошибки, если происходит деление на ноль
            }
            nums[i] /= nums[i+1]
            nums = append(nums[:i+1], nums[i+2:]...)
            operations = append(operations[:i], operations[i+1:]...)
            i--
        }
    }

    result := nums[0] // Инициализация результата первым числом

    // Выполнение сложения и вычитания
    for i, op := range operations {
        switch op {
        case "+":
            result += nums[i+1]
        case "-":
            result -= nums[i+1]
        }
    }

    return strconv.Itoa(result) // Возвращение результата в виде строки
}
