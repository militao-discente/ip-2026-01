package main

import "fmt"

func main() {
	var x, fx float64
	fmt.Print("Digite um valor para x: ")
	fmt.Scan(&x)
	fx = 8 / (2 - x)
	fmt.Printf("Se x = %.2f, então f(x) = %.2f\n", x, fx)
}
