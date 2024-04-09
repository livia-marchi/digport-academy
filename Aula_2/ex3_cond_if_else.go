package main

import "fmt"

func main() {

	var numero int

	fmt.Println("------ Exercicio 3 - Condicionais IF / ELSE -------")
	fmt.Printf("Digite um número: ")
	fmt.Scan(&numero)
	if numero > 0 {
		fmt.Println("O número é positivo")
	} else if numero < 0 {
		fmt.Println("O número é negativo")
	} else {
		fmt.Println("O número é zero")
	}
}
