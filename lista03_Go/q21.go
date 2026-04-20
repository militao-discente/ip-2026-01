package main

import "fmt"

func main() {
	var b, n int
	fmt.Print("Digite b (>=2): ")
	fmt.Scan(&b)
	fmt.Print("Digite n (>1): ")
	fmt.Scan(&n)
	resultado := 1
	for i := 0; i < n; i++ {
		resultado *= b
	}
	fmt.Printf("b^%d = %d\n", n, resultado)
}
