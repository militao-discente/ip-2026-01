package main

import "fmt"

func main() {
	var A, B int
	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&A)
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&B)
	if A%B == 0 {
		fmt.Printf("%d é divisível de %d\n", A, B)
	} else {
		fmt.Printf("%d não é divisível de %d\n", A, B)
	}
}
