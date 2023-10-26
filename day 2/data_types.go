// Create a Go program that declares and initializes variables of different data types.

package main

import "fmt"

func main() {
	// Integer variables
	var age int = 30
	var score int64 = 100

	// Floating-point variables
	var temp float32 = 72.5
	var pi float64 = 3.1416121626

	// Boolean variable
	var isTrue bool = true

	// String variable
	var name string = "John Doe"

	// Array
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	var names [5]string
	names[0] = "loki"
	names[1] = "gube"
	names[2] = "chinnu"
	names[3] = "maa"
	names[4] = "dummi"

	// Slice
	fruits := []string{"apple", "banana", "cherry"}
	fruits = append(fruits, "orange")

	fmt.Println("Age:", age)
	fmt.Println("Score:", score)
	fmt.Println("Temperature:", temp)
	fmt.Println("Pi:", pi)
	fmt.Println("Is True:", isTrue)
	fmt.Println("Name:", name)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Names:", names)
	fmt.Println("Fruits:", fruits)
}
