package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romeDict = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var allRoman = [][2]interface{}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"},
	{1, "I"},
}

// Функции математических операций
func sum(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func division(a, b int) int {
	return a / b
}

func mult(a, b int) int {
	return a * b
}

// Функция перевода чисел в римскую систему счисления
func toRoman(num int) string {
	roman := ""
	for num > 0 {
		for _, pair := range allRoman {
			value := pair[0].(int)
			romanNumeral := pair[1].(string)
			for num >= value {
				roman += romanNumeral
				num -= value
			}
		}
	}
	return roman
}

// Функция проверки ошибки. Выдаёт панику если ошибка
func checkError(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

// Запуск основной программы
func main() {
	// создаём reader указываем ввод с консоли
	reader := bufio.NewReader(os.Stdin)

	// Бесконечный цикл
	for {
		fmt.Println("Введите выражение (например: VI * IX):")
		inp, _ := reader.ReadString('\n') // получаем ввод из консоли
		inp = strings.TrimSpace(inp)

		parts := strings.Fields(inp) // Разбиваем input на части
		checkError(len(parts) == 3, "Ошибка: неверный ввод. Используйте формат 'число оператор число'.")

		num1, sign, num2 := parts[0], parts[1], parts[2]

		/* проверяем какие мы получили числа, римские или арабские
		Если римские, то ok1 и ok2 будут в значении true */
		n1, ok1 := romeDict[num1]
		n2, ok2 := romeDict[num2]

		// Если римские числа
		if ok1 && ok2 {
			checkError(n1 > 0 && n1 <= 10 && n2 > 0 && n2 <= 10, "Ошибка: числа должны быть от I до X.")

			// Выполняем вычисление
			switch sign {
			case "+":
				fmt.Println(toRoman(sum(n1, n2)))
			case "-":
				result := minus(n1, n2)
				checkError(result > 0, "Ошибка: результат отрицательный или ноль.")
				fmt.Println(toRoman(result))
			case "/":
				result := division(n1, n2)
				checkError(result >= 1, "Ошибка: результат отрицательный или ноль.")
				fmt.Println(toRoman(result))
			case "*":
				fmt.Println(toRoman(mult(n1, n2)))
			default:
				panic("Ошибка: неверный оператор.")
			}
			// Если арабские числа
		} else if !ok1 && !ok2 {
			// Переводим числа в int тип
			num1Int, err1 := strconv.Atoi(num1)
			num2Int, err2 := strconv.Atoi(num2)

			// Проверяем числа, если не перевелись то паника
			checkError(err1 == nil && err2 == nil, "Ошибка: неверный ввод чисел.")
			checkError(num1Int > 0 && num1Int <= 10 && num2Int > 0 && num2Int <= 10, "Ошибка: должны быть целые числа от 1 до 10.")

			// Выполняем математические операции
			switch sign {
			case "+":
				fmt.Println(sum(num1Int, num2Int))
			case "-":
				fmt.Println(minus(num1Int, num2Int))
			case "/":
				checkError(num2Int != 0, "Ошибка: деление на ноль.")
				fmt.Println(division(num1Int, num2Int))
			case "*":
				fmt.Println(mult(num1Int, num2Int))
			default:
				panic("Ошибка: неверный оператор.")
			}
			// Паника если невернный ввод чисел
		} else {
			panic("Ошибка: неверный ввод чисел.")
		}
	}
}
