package main

import (
	"fmt"
	"math"
)

func main() {
	s := 0.0
	for k := 0; k < 51; k++ {
		den := 2*k + 1
		termo := 1.0 / float64(den*den*den)
		if k%2 == 0 {
			s += termo
		} else {
			s -= termo
		}
	}
	pi := math.Pow(32*s, 1.0/3.0)
	fmt.Printf("Pi aproximado: %.10f\n", pi)
}
