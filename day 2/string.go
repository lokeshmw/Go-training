// Develop a Go program that demonstrates string manipulation
// by concatenating two strings and finding the length of the result.

package main

import "fmt"


func main(){
	var a, b string
	fmt.Println("Enter First string: ")
	fmt.Scanln(&a)
	fmt.Println("Enter Second string: ")
	fmt.Scanln(&b)
	con := a+b
	fmt.Println(con)
	fmt.Print("len: ", len(con))
}