package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Print(fact(0))
	fmt.Print(fact(1))
	fmt.Print(fact(2))
	fmt.Print(fact(3))

}
