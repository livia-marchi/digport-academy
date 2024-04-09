package main

import "fmt"

func main() {

	var numero int

	fmt.Println("------ Exercicio 3 - Condicionais SWITCH -------")
	fmt.Printf("Digite um número: ")
	fmt.Scan(&numero)

	switch {
	case numero > 0:
		fmt.Println("O número é positivo")
	case numero < 0:
		fmt.Println("O número é negativo")
	default:
		fmt.Println("O número é zero")
	}
}
