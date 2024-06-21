go
package main

import (
	"fmt"
)

func romanToArabic(roman string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prev := 0
	for i := len(roman) - 1; i >= 0; i-- {
		val := romanMap[rune(roman[i])]
		if val < prev {
			total -= val
		} else {
			total += val
		}
		prev = val
	}

	return total
}

func calculate(op string, a, b float64) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("деление на 0")
		}
		return a / b
	default:
		panic("неверный оператор")
	}
}

func main() {
	var operator, num1s, num2s string
	var num1, num2 float64
	var useRoman bool

	fmt.Print("Хотите ли использовать римские цифры? (yes/no): ")
	fmt.Scanln(&num1s)
	if num1s == "yes" {
		useRoman = true
	}

	if useRoman {
		fmt.Print("Введите число 1 в римских цифрах: ")
		fmt.Scanln(&num1s)
		num1 = float64(romanToArabic(num1s))

		fmt.Print("Введите число 2 в римских цифрах: ")
		fmt.Scanln(&num2s)
		num2 = float64(romanToArabic(num2s))
	} else {
		fmt.Print("Введите число 1: ")
		fmt.Scanln(&num1)

		fmt.Print("Введите число 2: ")
		fmt.Scanln(&num2)
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scanln(&operator)

	result := calculate(operator, num1, num2)
	fmt.Println("Результат:", result)
}	
