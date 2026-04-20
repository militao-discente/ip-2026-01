package main

import "fmt"

func main() {
	total := uint64(0)
	grao := uint64(1)
	for i := 1; i <= 64; i++ {
		total += grao
		grao *= 2
	}
	fmt.Printf("Total de graos: %d\n", total)
}
