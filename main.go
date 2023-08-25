package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Split(input, " ")
	if len(parts) > 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	} else if len(parts) < 3 {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	a, err := parseNumber(parts[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := parseNumber(parts[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := performOperation(a, b, parts[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if isRoman(parts[0]) && isRoman(parts[2]) {
		if result <= 0 {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел и нуля.")
			return
		}
		fmt.Println(toRoman(result))
	} else if !isRoman(parts[0]) && !isRoman(parts[2]) {
		fmt.Println(result)
	} else {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	}
}

// вычисление значения
func performOperation(a, b int, operation string) (int, error) {
	var result int
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Вывод ошибки, так как деление на ноль.")
		}
		result = a / b
	default:
		fmt.Println()
		return 0, fmt.Errorf("Вывод ошибки, так как неверная операция.")
	}
	return result, nil
}

// парсинг операнда
func parseNumber(s string) (int, error) {
	if isRoman(s) {
		return fromRoman(s)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Вывод ошибки, так как число представлено неверно.")
	}
	if n > 10 {
		return 0, fmt.Errorf("Вывод ошибки, так как число больше 10.")
	}
	return n, nil
}

// проверка на римскую цифру
func isRoman(s string) bool {
	symbol := []string {"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, value := range symbol {
		if s == value {
			return true
		}
	}
	return false
}

// перевод римской цифры в число
func fromRoman(s string) (int, error) {
	var result int
	for i := 0; i < len(s); i++ {
		if i > 0 && romanValue(s[i]) > romanValue(s[i-1]) {
			result += romanValue(s[i]) - 2*romanValue(s[i-1])
		} else {
			result += romanValue(s[i])
		}
	}
	if result > 10 {
		return 0, fmt.Errorf("Вывод ошибки, так как число больше 10.")
	}
	return result, nil
}

// преобразование результата в римскую цифру
func toRoman(n int) string {
	var result strings.Builder
	for _, r := range []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{500, "D"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},	
	} {
		for n >= r.value {
			result.WriteString(r.symbol)
			n -= r.value
		}
	}
	return result.String()
}

// соотвествие римской цифры числу
func romanValue(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	default:
		return 0
	}
}
