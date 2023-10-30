// Create a Go program that calculates the simple interest on a loan.
// Prompt the user for the principal amount, interest rate, and time (in years).
// Implement error handling for invalid input values.

package main

import "fmt"

func main(){
	var p, r, t float64
	fmt.Println("Enter the Principal amount: ")
	fmt.Scanln(&p)
	fmt.Println("Enter the Rate of interst: ")
	fmt.Scanln(&r)
	fmt.Println("Enter the Time in years: ")
	fmt.Scanln(&t)

	if p < 0 || r < 0 || t < 0 {
		fmt.Println("Please Enter vaild details It can't be Negative")
	}else{
	interest := (p * r * t) / 100
	fmt.Println("Simple interst: ", interest)
	}
}