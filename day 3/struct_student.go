package main

import (
	"fmt"
)

type student struct {
	name       string
	id         int
	class      int
	percentage float64
	grade      string
}

func main() {
	lokesh := student{
		name:       "Lokesh",
		id:         423,
		class:      7,
		percentage: 99,
	}
	Gube := student{
		name:       "usisiyaa",
		id:         420,
		class:      5,
		percentage: 99.99,
	}
	fmt.Print(lokesh.g())
	fmt.Print(Gube.g())

}

func (a *student) g() string {
	if a.percentage < 35 {
		a.grade = "D"
		return a.grade
	}if a.percentage >= 35 || a.percentage <= 50 {
		a.grade = "C"
		return a.grade
	} else if a.percentage >= 50 || a.percentage <= 80 {
		a.grade = "B"
		return a.grade
	}else{
	a.grade = "A"
	return a.grade
	}
}
