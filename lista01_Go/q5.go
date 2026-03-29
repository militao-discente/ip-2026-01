package main

import "fmt"

func main() {
	var conta int
	var consumo float64
	var tipo string
	fmt.Scan(&conta, &consumo, &tipo)
	var valor float64
	switch tipo {
	case "R", "r":
		valor = 5.0 + 0.05*consumo
	case "C", "c":
		valor = 500.0
		if consumo > 80.0 {
			valor += 0.25 * (consumo - 80.0)
		}
	case "I", "i":
		valor = 800.0
		if consumo > 100.0 {
			valor += 0.04 * (consumo - 100.0)
		}
	}
	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}
