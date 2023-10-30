// Write a Go program that calculates the area and circumference of a circle.
//  Prompt the user for the radius and handle errors for non-numeric or negative inputs.

package main

import (
	"fmt"
	"math")

func a(r float64) float64 {
	if r < 0 {
		fmt.Println("Enter a positive number")
		return 0
	}
	if r == 0 {
		return 0
	}
	{
		return math.Pi * math.Pow(r, 2)
	}
}

func c(r float64) float64 {
	if r < 0 {
		return 0
	}
	if r == 0 {
		return 0
	}
	{
		return math.Pi * 2 * r
	}
}

func main() {
	var radius float64
	fmt.Println("Enter raidus: ")
	fmt.Scanln(&radius)
	fmt.Println("Area of circle: ", a(radius))
	fmt.Println("Circumference of circle: ", c(radius))
}
