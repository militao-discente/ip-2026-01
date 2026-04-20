package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)
	produto := 0
	for i := 0; i < n2; i++ {
		produto += n1
	}
	fmt.Printf("%d * %d = %d\n", n1, n2, produto)
}
