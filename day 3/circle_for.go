package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Println("Enter raidus: ")
	fmt.Scanln(&radius)
	if radius < 0 {
		fmt.Println("Enter a positive number")
	}
	if radius == 0 {
		fmt.Println("Enter a positive number greater than 0 ")

	}
	if radius > 0 {
		area := math.Pi * math.Pow(radius, 2)
		circumference := math.Pi * 2 * radius
		fmt.Println("Area of circle: ", area)
		fmt.Println("Circumference of circle: ", circumference)
	}
}
