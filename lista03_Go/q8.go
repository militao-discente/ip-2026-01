package main

import "fmt"

func main() {
	soma := 0
	for i := 1; i <= 20; i++ {
		fmt.Printf("%d ", i)
		soma += i
	}
	fmt.Printf("\nSoma: %d\n", soma)
}
