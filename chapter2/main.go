package main

import "fmt"

func main() {
	// Declare variables to store user input
	var num1, num2 int

	// Prompt user to enter the first number and read input
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	// Prompt user to enter the second number and read input
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	// Perform arithmetic operations
	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	// Display the results
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)
}
