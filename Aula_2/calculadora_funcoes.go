package main

import "fmt"

func numbers() (float32, float32, float32, float32) {
	var a, b float32
	fmt.Printf("\nEnter the first number: ")
	fmt.Scan(&a)
	fmt.Printf("Enter the second number: ")
	fmt.Scan(&b)
	sum := a + b
	sub := a - b
	mult := a * b
	div := a / b
	return sum, sub, mult, div
}

func one_more_time(op int) int {
	fmt.Println("\nIf you want to use the calculator again, press 0. To exit, press any key.")
	fmt.Scan(&op)
	if op == 0 {
		return op
	} else {
		fmt.Println("\nThank you for using this calculator. Stay tuned for the next projects!\n")
		return op
	}
}

func main() {

	var op int
	op = 0

	//for op != 1 && op != 2 && op != 3 && op != 4 && op != 5 {
	for {
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
			sum, _, _, _ := numbers()
			fmt.Println("The result of the sum of the values is: ", sum)
			op = one_more_time(op)
		case 2:
			_, sub, _, _ := numbers()
			fmt.Println("The result of the subtraction of the values is: ", sub)
			op = one_more_time(op)
		case 3:
			_, _, mult, _ := numbers()
			fmt.Println("The result of the multiplication of the values is: ", mult)
			op = one_more_time(op)
		case 4:
			_, _, _, div := numbers()
			fmt.Println("The result of the division of the values is: ", div)
			op = one_more_time(op)
		case 5:
			fmt.Println("\nThank you for using this calculator. Stay tuned for the next projects!\n")
		default:
			fmt.Println("\nInvalid value. Please, try again.\n")
		}
	}
}
