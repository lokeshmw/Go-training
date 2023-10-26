// Write a Go program that performs basic arithmetic operations (addition, subtraction, multiplication, and division)
//  on two numbers.

package main

import "fmt"

func main() {
	var a, b int
	var choice int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&b)
	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("Enter your choice (1/2/3/4): ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("sum: ", a+b)

	case 2:
		fmt.Println("diff: ", a-b)

	case 3:
		fmt.Println("product: ", a*b)
	case 4:
		fmt.Println("div: ", a/b)
	default:
		fmt.Println("Invalid choice")

	}

}
