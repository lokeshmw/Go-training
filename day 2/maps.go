package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	i, prs := m["k2"]
	fmt.Println("p:", i)
	fmt.Println("prs:", prs)

	j, prrs := m["k3"]
	fmt.Println("p:", j)
	fmt.Println("prs:", prrs)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	v4 := m["k4"]
	fmt.Println("v4:", v4)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2, "a": 45}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2, "a": 45}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
