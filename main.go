package main

import "fmt"

func main() {
	var (
		expression         string
		num1, num2, result float64
		operation          rune
	)
	for {
		fmt.Print("Введи выражение (или exit для выхода): ")
		fmt.Scanln(&expression)
		if expression == "exit" {
			fmt.Println("Программа заверщена!")
			break
		}
		_, err := fmt.Sscanf(expression, "%f%c%f", &num1, &operation, &num2)
		if err != nil {
			fmt.Println("Ошибка при считывания данных")
			continue
		}
		switch operation {
		case '+':
			result = num1 + num2
		case '-':
			result = num1 - num2
		case '/':
			if num2 == 0 {
				fmt.Println("Ошибка, на НОЛЬ делить нельзя")
				continue
			}
			result = num1 / num2
		case '*':
			result = num1 * num2
		}
		fmt.Println("Результат: ", result)

	}
}
