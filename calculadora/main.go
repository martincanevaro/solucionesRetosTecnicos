package main

import (
	"fmt"

	"./calc"
)

func main() {
	// Crear una instancia de la calculadora
	calculator := calc.NewCalculator()

	// Pedir al usuario los valores
	fmt.Print("Introduce el primer valor: ")
	a, _ := fmt.Scanf("%d", &a)
	fmt.Print("Introduce el segundo valor: ")
	b, _ := fmt.Scanf("%d", &b)

	// Pedir al usuario el cálculo que desea hacer
	fmt.Print("¿Qué cálculo quieres realizar? (suma, resta, multiplicación, división): ")
	operation, _ := fmt.Scanf("%s", &operation)

	// Realizar el cálculo
	switch operation {
	case "suma":
		fmt.Println(calculator.Add(a, b))
	case "resta":
		fmt.Println(calculator.Subtract(a, b))
	case "multiplicación":
		fmt.Println(calculator.Multiply(a, b))
	case "división":
		fmt.Println(calculator.Divide(a, b))
	default:
		fmt.Println("Operación no válida")
	}
}
