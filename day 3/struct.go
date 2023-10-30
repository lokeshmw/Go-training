package main

import "fmt"

type Bank struct {
	Acc_no         int
	name           string
	balance        float64
	rate_interst   float64
	account_age    float64
	interset_value float64
}

func main() {
	Lokesh_acc := Bank{
		Acc_no:       45,
		name:         "Lokesh",
		balance:      1000,
		rate_interst: 9,
		account_age:  3,
	}
	fmt.Println(Lokesh_acc.Simple_interest())
}

func (A *Bank) Simple_interest() (float64, float64) {
	interset := A.balance * A.account_age * A.rate_interst / 100
	A.interset_value += interset
	A.balance += interset
	return A.interset_value, A.balance
}
