package main

import "fmt"

func a() func() int {
	i := 0
	return func() int {
		i++
		return i

	}
}

func main() {
	b:= a()
	fmt.Print(b())
	fmt.Print(b())
	fmt.Print(b())
	fmt.Print(b())
	fmt.Print(b())
	fmt.Print(b())
	fmt.Print(b())
	c:=a()
	fmt.Print(c())
	

}
