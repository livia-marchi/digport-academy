package main

import "fmt"

func main() {

	var nome string

	fmt.Println("------ Exercicio 1 - Bem-vindo -------")
	fmt.Print("Digite o seu nome: ")
	fmt.Scan(&nome)
	fmt.Println("Bem-vindo, " + nome + "!")
}
