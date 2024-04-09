package main

import "fmt"

func main() {

	const a = 1
	//var a int
	var b int
	var soma int

	fmt.Println("------ Exercicio 2 - Soma de 2 números -------")
	//fmt.Printf("Digite o primeiro número: \n")
	//fmt.Scan(&a)
	fmt.Printf("Digite o segundo número: \n")
	fmt.Scan(&b)
	soma = a + b
	fmt.Println("A soma é ", soma)
}
