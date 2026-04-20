package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite N: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Valor invalido")
		return
	}
	s := 0
	num := 1000
	den := 1
	for i := 1; i <= n; i++ {
		termo := num / den
		if i%2 == 1 {
			s += termo
		} else {
			s -= termo
		}
		num -= 3
		den++
	}
	fmt.Printf("Resultado = %d\n", s)
}
