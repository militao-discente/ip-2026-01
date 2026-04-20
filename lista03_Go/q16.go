package main

import "fmt"

func main() {
	var n, a1, a2 int
	fmt.Print("Digite N (>=3): ")
	fmt.Scan(&n)
	fmt.Print("Digite A1: ")
	fmt.Scan(&a1)
	fmt.Print("Digite A2: ")
	fmt.Scan(&a2)
	fmt.Printf("%d %d ", a1, a2)
	for i := 3; i <= n; i++ {
		var ai int
		if i%2 == 1 {
			ai = a2 + a1
		} else {
			ai = a2 - a1
		}
		fmt.Printf("%d ", ai)
		a1 = a2
		a2 = ai
	}
	fmt.Println()
}
