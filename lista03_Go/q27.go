package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Digite x: ")
	fmt.Scan(&x)
	cosCalc := 1.0
	fat := 1.0
	for i := 1; i <= 10; i++ {
		fat *= float64(2 * i * (2*i - 1))
		termo := math.Pow(x, float64(2*i)) / fat
		if i%2 == 1 {
			cosCalc -= termo
		} else {
			cosCalc += termo
		}
	}
	cosLib := math.Cos(x)
	fmt.Printf("Cos calculado: %.10f\n", cosCalc)
	fmt.Printf("Cos biblioteca: %.10f\n", cosLib)
	fmt.Printf("Diferenca: %.10f\n", cosCalc-cosLib)
}
