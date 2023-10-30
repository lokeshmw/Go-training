package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	res := add(11, 12)
	fmt.Print(res)

}
