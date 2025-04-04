package main

import (
	"fmt"
	"github.com/matheus-alvs01dev/eng-software/calculator"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Resultado: %.2f\n", calculator.Add(num1, num2))
	case "-":
		fmt.Printf("Resultado: %.2f\n", calculator.Subtract(num1, num2))
	case "*":
		fmt.Printf("Resultado: %.2f\n", calculator.Multiply(num1, num2))
	case "/":
		result, err := calculator.Divide(num1, num2)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Printf("Resultado: %.2f\n", result)
		}
	default:
		fmt.Println("Operador inválido. Use +, -, * ou /.")
	}
}
