package agent

import (
	"strconv"
	"strings"
	"time"
)

// Функция для вычисления результата выражения и обновления статуса и значения
func ProcessExpression(expression string, tim int) string {
	t := time.Duration(tim) * time.Second

	resultChan := make(chan string, 1)

	go func() {
		time.Sleep(t)
		resultChan <- calculate(expression)
	}()

	return <-resultChan
}

func calculate(expression string) string {
	tokens := strings.FieldsFunc(expression, func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	})

	nums := make([]int, 0)
	operations := make([]string, 0)

	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			return "Ошибка: неверный формат числа"
		}
		nums = append(nums, num)
	}

	for _, c := range expression {
		char := string(c)
		if char == "+" || char == "-" || char == "*" || char == "/" {
			operations = append(operations, char)
		}
	}

	for i := 0; i < len(operations); i++ {
		if operations[i] == "*" {
			nums[i] *= nums[i+1]
			nums = append(nums[:i+1], nums[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--
		} else if operations[i] == "/" {
			if nums[i+1] == 0 {
				return "Ошибка: деление на ноль"
			}
			nums[i] /= nums[i+1]
			nums = append(nums[:i+1], nums[i+2:]...)
			operations = append(operations[:i], operations[i+1:]...)
			i--
		}
	}

	result := nums[0]

	for i, op := range operations {
		switch op {
		case "+":
			result += nums[i+1]
		case "-":
			result -= nums[i+1]
		}
	}

	return strconv.Itoa(result)
}
