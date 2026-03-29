package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		var publico int
		var pp, pg, pa, pc float64
		fmt.Scan(&publico, &pp, &pg, &pa, &pc)
		pop := float64(publico) * pp / 100.0
		ger := float64(publico) * pg / 100.0
		arb := float64(publico) * pa / 100.0
		cad := float64(publico) * pc / 100.0
		renda := pop*1.0 + ger*5.0 + arb*10.0 + cad*20.0
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
