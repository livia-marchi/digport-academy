package main

import "fmt"

func main() {
	lista := map[string]string{"Ana": "escalar", "Bruno": "desenhar", "Cecília": "escrever", "Diogo": "nadar"}
	//Para acessar o hobby de uma pessoa da lista
	fmt.Println("O hobby de Ana é", lista["Ana"])
	//Para acessar a lista de todas as pessoas e seus hobbies
	for key, value := range lista {
		fmt.Printf("O hobby de %s é %s\n", key, value)
	}
}
