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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: VI * IX):")
	inp, _ := reader.ReadString('\n')
	inp = strings.TrimSpace(inp)

	parts := strings.Fields(inp)
	if len(parts) != 3 {
		fmt.Println("Ошибка: неверный ввод. Используйте формат 'число оператор число'.")
		return
	}

	num1, sign, num2 := parts[0], parts[1], parts[2]

	n1, ok1 := romeDict[num1]
	n2, ok2 := romeDict[num2]

	if ok1 && ok2 {
		if n1 > 0 && n1 <= 10 && n2 > 0 && n2 <= 10 {
			switch sign {
			case "+":
				fmt.Println(toRoman(sum(n1, n2)))
			case "-":
				result := minus(n1, n2)
				if result > 0 {
					fmt.Println(toRoman(result))
				} else {
					fmt.Println("Ошибка: результат отрицательный или ноль.")
				}
			case "/":
				if n2 != 0 {
					fmt.Println(toRoman(division(n1, n2)))
				} else {
					fmt.Println("Ошибка: деление на ноль.")
				}
			case "*":
				fmt.Println(toRoman(mult(n1, n2)))
			default:
				fmt.Println("Ошибка: неверный оператор.")
			}
		} else {
			fmt.Println("Ошибка: числа должны быть от I до X.")
		}
	} else if !ok1 && !ok2 {
		num1, _ := strconv.Atoi(num1)
		num2, _ := strconv.Atoi(num2)

		if num1 > 0 && num1 <= 10 && num2 > 0 && num2 <= 10 {
			switch sign {
			case "+":
				fmt.Println(sum(num1, num2))
			case "-":
				fmt.Println(minus(num1, num2))
			case "/":
				if num2 != 0 {
					fmt.Println(division(num1, num2))
				} else {
					fmt.Println("Ошибка: деление на ноль.")
				}
			case "*":
				fmt.Println(mult(num1, num2))
			default:
				fmt.Println("Ошибка: неверный оператор.")
			}
		} else {
			fmt.Println("Ошибка: должны быть целые числа от 1 до 10.")
		}
	} else {
		fmt.Println("Ошибка: неверный ввод чисел.")
	}
}
