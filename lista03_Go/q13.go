package main

import "fmt"

func main() {
	h := 0.0
	for i := 1; i <= 50; i++ {
		num := 2*i - 1
		h += float64(num) / float64(i)
	}
	fmt.Printf("H = %.10f\n", h)
}
