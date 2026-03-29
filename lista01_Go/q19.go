package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}
	s := 0.0
	for k := 1; k <= n; k++ {
		s += 1.0 / float64(k)
	}
	fmt.Printf("%.6f\n", s)
}
