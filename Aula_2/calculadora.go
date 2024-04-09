package main

import "fmt"

func main() {

	var num1, num2 float32
	var sum, sub, mult, div float32
	var cont int

	op := 0

	for op != 1 && op != 2 && op != 3 && op != 4 && op != 5 {
		fmt.Println("------------------------- Dig Calculator -------------------------")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Println("------------------------------------------------------------------\n")
		fmt.Printf("Please, choose the operation you want to execute: ")
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Printf("\nEnter the first number: ")
			fmt.Scan(&num1)
			fmt.Printf("Enter the second number: ")
			fmt.Scan(&num2)
			sum = num1 + num2
			fmt.Println("The result of the sum of the values is: ", sum)
			fmt.Println("\nDo you want to use the calculator again? 1.YES / 2. NO")
			fmt.Scan(&cont)
			if cont == 1 {
				op = 0
			} else {
				fmt.Println("\nThank you for using this calculator. Stay tuned for the next projects!\n")
			}
		case 2:
			fmt.Printf("\nEnter the first number: ")
			fmt.Scan(&num1)
			fmt.Printf("Enter the second number: ")
			fmt.Scan(&num2)
			sub = num1 - num2
			fmt.Println("The result of the subtraction of the values is: ", sub)
			fmt.Println("\nDo you want to use the calculator again? 1.YES / 2. NO")
			fmt.Scan(&cont)
			if cont == 1 {
				op = 0
			} else {
				fmt.Println("\nThank you for using this calculator. Stay tuned for the next projects!\n")
			}
		case 3:
			fmt.Printf("\nEnter the first number: ")
			fmt.Scan(&num1)
			fmt.Printf("Enter the second number: ")
			fmt.Scan(&num2)
			mult = num1 * num2
			fmt.Println("The result of the multiplication of the values is: ", mult)
			fmt.Println("\nDo you want to use the calculator again? 1.YES / 2. NO")
			fmt.Scan(&cont)
			if cont == 1 {
				op = 0
			} else {
				fmt.Println("\nThank you for using this calculator. Stay tuned for the next projects!\n")
			}
		case 4:
			fmt.Printf("\nEnter the first number: ")
			fmt.Scan(&num1)
			fmt.Printf("Enter the second number: ")
			fmt.Scan(&num2)
			div = num1 / num2
			fmt.Println("The result of the division of the values is: ", div)
			fmt.Println("\nDo you want to use the calculator again? 1.YES / 2. NO")
			fmt.Scan(&cont)
			if cont == 1 {
				op = 0
			} else {
				fmt.Println("\nThank you for using this calculator. Stay tuned for the next projects!\n")
			}
		case 5:
			fmt.Println("\nThank you for using this calculator. Stay tuned for the next projects!\n")
		default:
			fmt.Println("\nInvalid value. Please, try again.\n")
		}
	}
}
