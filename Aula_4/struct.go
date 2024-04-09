// 1. Crie uma structure para um funcionário onde deseja-se saber nome, função, idade e salario
// 2. Atributa pelo menos 3 funcionários - e.g: Lorraine, professora, 35, 3000
// 3. Imprima os valores da segunda pessoa.

package main

import "fmt"

type Funcionario struct {
	nome    string
	funcao  string
	idade   int
	salario float32
}

func main() {

	var funA Funcionario
	var funB Funcionario
	var funC Funcionario

	funA.nome = "Tony"
	funA.funcao = "Engenheiro"
	funA.idade = 47
	funA.salario = 50000.0

	funB.nome = "Marie"
	funB.funcao = "Quimica"
	funB.idade = 44
	funB.salario = 5000.0

	funC.nome = "Peter"
	funC.funcao = "Cientista"
	funC.idade = 20
	funC.salario = 2000.0

	fmt.Println(funB.nome)
	fmt.Println(funB.funcao)
	fmt.Println(funB.idade)
	fmt.Println(funB.salario)

}
