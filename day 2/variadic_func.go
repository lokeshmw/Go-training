package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total
}

func concatenat(names ...string) string{
	ans:=""
	for _, j:=range names{
		ans+=j
	}
	return ans

}

func main() {
	fmt.Println(sum(12, 1))
	fmt.Println(sum(1, 23, 2))
	fmt.Print(concatenat("Lokesh", "M"))

}

