package main

import "fmt"

func main() {
	soma := 0
	countador := 0
	for i := 50; i <= 70; i += 2 {
		soma += i
		countador++
	}
	media := float64(soma) / float64(countador)
	fmt.Printf("Soma dos pares: %d\nMédia aritmética: %.2f\n", soma, media)
}
