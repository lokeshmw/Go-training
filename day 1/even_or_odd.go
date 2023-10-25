// In a given range split odd and even numbers and printing them individually
package main

import "fmt"

func main(){
	var first_limit int
	var last_limit int
	fmt.Println("Enter First Limit: ")
	fmt.Scanln(&first_limit)
	fmt.Println("Enter Last Limit: ")
	fmt.Scanln(&last_limit)
	var even, odd   []int
	for i:= first_limit; i <= last_limit; i++{
		if i % 2 == 0{
			even = append(even, i)
			} else{
			odd = append(odd, i)
			}
		}
	fmt.Println(even)
	fmt.Println(odd)


}