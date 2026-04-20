package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1: ")
	fmt.Scan(&n1)
	fmt.Print("Digite N2: ")
	fmt.Scan(&n2)
	a := n1
	b := n2
	for b != 0 {
		a, b = b, a%b
	}
	lcm := n1 * n2 / a
	fmt.Printf("MMC(%d,%d) = %d\n", n1, n2, lcm)
}
