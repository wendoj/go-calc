package main

import "fmt"

func calculate(operator string, a, b int) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if a == 0 || b == 0 {
			fmt.Println("Cannot divide by zero")
			return 0
		}
		return a / b
	case "%":
		if b == 0 {
			fmt.Println("Cannot divide by zero")
			return 0
		}
		return a % b
	default:
		fmt.Println("Operator not supported")
		return 0
	}
}

func main() {
	/* Get user input */
	var operator string
	var a, b int

	fmt.Println("Enter operator:")
	fmt.Scanln(&operator)
	fmt.Println("Enter two numbers (separate with a space):")
	fmt.Scanln(&a, &b)

	/* Calculate and print result */
	fmt.Println(calculate(operator, a, b))
}
