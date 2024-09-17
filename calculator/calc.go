package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var deist, kon string
	for kon != "n" {
		fmt.Println("Введите числа для калькулятора: ")
		fmt.Scanln(&a)
		fmt.Scanln(&b)
		fmt.Print("Выберите действие: \n 1. Сложение (+) \n 2. Вычитание (-) \n 3. Деление (/) \n 4. Умножение (*) \n")
		fmt.Scanln(&deist)
		switch deist {
		case "+":
			fmt.Println("Результат: ", a+b)
		case "-":
			fmt.Println("Результат: ", a-b)
		case "/":
			if b != 0 {
				fmt.Println("Результат: ", a/b)
			} else {
				fmt.Println("Деление на 0 невозможно")
				break
			}
		case "*":
			fmt.Println("Результат: ", a*b)
		}
		fmt.Println("Продолжить?(n/y) ")
		fmt.Scanln(&kon)
	}
}
