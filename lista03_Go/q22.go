package main

import "fmt"

func main() {
	s := 0
	for i := 1; i <= 37; i++ {
		termo := (38 - i) * (39 - i) / i
		s += termo
	}
	fmt.Printf("S = %d\n", s)
}
