package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	var deist, kon string
	fmt.Println("Введите числа для калькулятора: ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Print("Выберите действие: \n 1. Сложение (+) \n 2. Вычитание (-) \n 3. Деление (/) \n 4. Умножение (*) \n")
	fmt.Scanln(&deist)
	for i := 0; i < 1000; i++ {
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
		if kon == "n" {
			break
		}
	}
}
