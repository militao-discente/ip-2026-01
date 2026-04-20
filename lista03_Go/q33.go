package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)
	quociente := 0
	resto := n1
	for resto >= n2 {
		resto -= n2
		quociente++
	}
	fmt.Printf("Quociente: %d\n", quociente)
	fmt.Printf("Resto: %d\n", resto)
}
