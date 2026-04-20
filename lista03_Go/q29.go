package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite N: ")
	fmt.Scan(&n)
	soma := n * (n + 1) / 2
	fmt.Printf("Soma de 1 a %d = %d\n", n, soma)
}
