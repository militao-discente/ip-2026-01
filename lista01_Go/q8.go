package main

import "fmt"

func main() {
	var r, a float64
	fmt.Scan(&r)
	fmt.Scan(&a)
	pi := 3.14159
	ac := pi * r * r
	al := 2.0 * pi * r * a
	at := 2.0*ac + al
	custo := at * 100.0
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
