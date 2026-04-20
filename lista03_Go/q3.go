package main

import "fmt"

func main() {
	var carlos float64
	fmt.Scan(&carlos)
	joao := carlos / 3.0

	r := 0
	for joao < carlos {
		carlos += carlos * 0.02
		joao += joao * 0.05
		r++
	}
	fmt.Printf("%d Mêses", r)
}
