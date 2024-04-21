package agent

import (
	"testing"
)

func TestProcessExpression_Success(t *testing.T) {
	tests := []struct {
		expression string
		tim        int
		expected   string
	}{
		{"2 + 2", 1, "4"},
		{"10 - 5", 2, "5"},
		{"3 * 4", 3, "12"},
		{"12 / 3", 4, "4"},
		{"2 + 3 * 4", 5, "14"}, // Проверка приоритета операций
	}

	for _, test := range tests {
		result := ProcessExpression(test.expression, test.tim)
		if result != test.expected {
			t.Errorf("ProcessExpression(%q, %d) = %q, expected %q", test.expression, test.tim, result, test.expected)
		}
	}
}

func TestProcessExpression_Error(t *testing.T) {
	tests := []struct {
		expression string
		tim        int
		expected   string
	}{
		{"2 + a", 1, "Ошибка: неверный формат числа"},
		{"10 / 0", 2, "Ошибка: деление на ноль"},
	}

	for _, test := range tests {
		result := ProcessExpression(test.expression, test.tim)
		if result != test.expected {
			t.Errorf("ProcessExpression(%q, %d) = %q, expected %q", test.expression, test.tim, result, test.expected)
		}
	}
}

