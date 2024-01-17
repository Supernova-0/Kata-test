package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для конвертации арабских чисел в римские
func arabicToRoman(n int) string {
	if n < 1 || n > 3999 {
		return ""
	}
	roman := ""
	arabic := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i, a := range arabic {
		for n >= a {
			roman += symbols[i]
			n -= a
		}
	}
	return roman
}

// Функция для конвертации римских чисел в арабские
func romanToArabic(s string) int {
	numbers := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	if s == "" {
		return 0
	}
	a, ok := numbers[s]
	if !ok {
		return 0
	}
	return a
}

// Функция для проверки, является ли строка римским числом
func isRoman(s string) bool {
	symbols := map[byte]bool{'I': true, 'V': true, 'X': true, 'L': true, 'C': true}
	for i := 0; i < len(s); i++ {
		if !symbols[s[i]] {
			return false
		}
	}
	return true
}

// Функция для проверки, является ли строка арабским числом
func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Функция для проверки, является ли строка операцией
func isOperation(s string) bool {
	operations := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	return operations[s]
}

// Функция для выполнения операции над двумя числами
func operate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль")
		}
		return a / b
	default:
		panic("Неверная операция")
	}
}

// Функция для обработки входной строки и вывода результата
func process(input string) {
	input = strings.ReplaceAll(input, " ", "")

	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	})
	if len(parts) != 2 {
		panic("Неверный формат ввода")
	}

	roman := isRoman(parts[0]) && isRoman(parts[1])
	arabic := isArabic(parts[0]) && isArabic(parts[1])
	if !roman && !arabic {
		panic("Числа должны быть одновременно арабскими или римскими")
	}

	a, b := parts[0], parts[1]
	if roman {
		a = strconv.Itoa(romanToArabic(parts[0]))
		b = strconv.Itoa(romanToArabic(parts[1]))
	}

	x, err1 := strconv.Atoi(a)
	y, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil || x < 1 || x > 10 || y < 1 || y > 10 {
		panic("Числа должны быть от 1 до 10 включительно")
	}

	op := ""
	for _, r := range input {
		if isOperation(string(r)) {
			op = string(r)
			break
		}
	}
	if op == "" {
		panic("Нет операции в строке")
	}

	result := operate(x, y, op)
	if roman {
		if result < 1 {
			panic("Результат работы калькулятора c римскими числами должен быть положителен")
		}
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()

	process(input)
}
