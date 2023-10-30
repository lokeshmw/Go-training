package main

import (
	"fmt"
	"errors"
)

var Diverror = errors.New("failed")

func divide(a,b int)(int, error){
	if b == 0{
		return -1, Diverror
	}
	return a/b, nil

}
func main(){
	var i, j int
	fmt.Println("ENter two numbers: ")
	fmt.Scanln(&i, &j)
	val, err :=divide(i, j)
	if err != nil{
		fmt.Print(err)
	  }
	fmt.Print(val)
}

