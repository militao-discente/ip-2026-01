package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	for i := n1; i <= n2; i++ {
		ehPrimo := true
		if i <= 1 {
			ehPrimo = false
		}
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				ehPrimo = false
				break
			}
		}
		if ehPrimo {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
