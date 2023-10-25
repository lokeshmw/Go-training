// Create a Go program that uses control structures (if statements) to check if a number is prime and odd/even.
package main

import "fmt"

func main(){
	var num int
	fmt.Println("Enter a number ")
	fmt.Scanln(&num)
	if num == 0{
		fmt.Println("num is not prime and even")
	}
	if num == 1{
		fmt.Println("num is prime and odd")

	}
	if num>=2{
		for i:= 2; i <= num; i++{
			if num % i == 0 && num%2 == 0{
				fmt.Println("num is prime and even")
				}else if  num % i == 0 && num%2 != 0{
					fmt.Println("num is prime and odd")
			}
		}
}
}