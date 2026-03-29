package main

import "fmt"

func main() {
	var salario float64
	fmt.Scan(&salario)
	var novo float64
	if salario <= 300.0 {
		novo = salario * 1.5
	} else {
		novo = salario * 1.3
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", novo)
}
