package main

import "fmt"

func main() {
	agenda := map[string]int{"Ana": 22222222, "Bruno": 33333333, "Cecília": 44444444, "Diogo": 55555555}
	//Para acessar toda a agenda
	fmt.Println(agenda)
	//Para acessar um item específico da agenda
	fmt.Println(agenda["Diogo"])
}
