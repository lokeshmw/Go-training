package main

import (
	"errors"
	"fmt"
	"math"
)

var CircleErr = errors.New("Enter vaild input")

func main() {
	var radius float64
	fmt.Println("Enter raidus: ")
	fmt.Scanln(&radius)
	result_area, error := area(radius)
	if error != nil {
		fmt.Println("failed:", error)
		return
	}
	fmt.Println("Result = ", result_area)

	result_circumference, error := circumference(radius)
	if error != nil {
		fmt.Println("failed:", error)
		return
	}
	fmt.Println("Result = ", result_circumference)
}
func area(radius float64) (float64, error) {
	if radius < 0 {
		return -1, CircleErr
	} else {
		Area_of_circle:= math.Pi * math.Pow(radius, 2)
		
		return Area_of_circle, nil
	}
}
func circumference(radius float64) (float64, error) {
	if radius < 0 {
		return -1, CircleErr
	} else {
		circumference_of_circle := math.Pi * 2 * radius
		return circumference_of_circle, nil
	}
}