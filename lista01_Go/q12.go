package main

import "fmt"

func main() {
	var horas int
	fmt.Scan(&horas)

	blocos := horas / 3
	restantes := horas % 3
	valor := float64(blocos)*10.0 + float64(restantes)*5.0

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}
