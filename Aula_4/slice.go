package main

import "fmt"

func main() {

	// 1. Crie uma constante com o valor “Hello”
	const a string = "Hello"

	// 2. Crie um slice de strings
	names := make([]string, 0, 3)

	// 3. Adicione ao slice: Hello (da constante), Nome1, Nome2, Nome3
	names = append(names, a, "Rivaldo", "Reinaldo", "Reginaldo")

	// 4. Atribua o segundo valor do slice a uma variável
	b := names[2]

	// 5. Imprima o slice e a frase “Hello, variavel”
	fmt.Println(names[0], ",", b)
}
