package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, j := range nums {
		sum += j
	}
	fmt.Println(sum)

}