// // Create a Go program that calculates the area of a rectangle. It should prompt the user for the length and width.

package main

import "fmt"

func main() {
	var len, wid int
	fmt.Println("Enter Len of rectange: ")
	fmt.Scanln(&len)
	fmt.Println("Enter Wid of rectange: ")
	fmt.Scanln(&wid)
	fmt.Println("Area of rectangle: ", len * wid)
}
