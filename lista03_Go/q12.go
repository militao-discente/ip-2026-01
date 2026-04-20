package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite X: ")
	fmt.Scan(&x)
	s := x
	fat := 1.0
	for k := 1; k < 20; k++ {
		fat *= float64(k)
		termo := x / fat
		if k%2 == 1 {
			s -= termo
		} else {
			s += termo
		}
	}
	fmt.Printf("S = %.10f\n", s)
}
