package main

import "fmt"

func main() {
	s := 0.0
	num := 1
	den := 225
	for i := 0; i < 9; i++ {
		termo := float64(num) / float64(den)
		if i%2 == 0 {
			s += termo
		} else {
			s -= termo
		}
		num *= 2
		den -= 27
	}
	fmt.Printf("S = %.10f\n", s)
}
