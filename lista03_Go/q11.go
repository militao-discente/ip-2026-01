package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite N: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Valor invalido")
		return
	}
	fat := 1
	for i := 2; i <= n; i++ {
		fat *= i
	}
	fmt.Printf("Fatorial de %d = %d\n", n, fat)
}
