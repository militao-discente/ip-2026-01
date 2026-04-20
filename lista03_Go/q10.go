package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite N (>=3): ")
	fmt.Scan(&n)
	a := 0
	b := 1
	fmt.Printf("%d %d ", a, b)
	for i := 2; i < n; i++ {
		c := a + b
		fmt.Printf("%d ", c)
		a = b
		b = c
	}
	fmt.Println()
}
