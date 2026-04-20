package main

import "fmt"

func main() {
	s := 0.0
	fat := 1.0
	for i := 0; i < 20; i++ {
		termo := float64(100-i) / fat
		s += termo
		fat *= float64(i + 1)
	}
	fmt.Printf("Soma = %.10f\n", s)
}
