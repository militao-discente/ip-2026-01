package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Digite o valor de a: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de b: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de c: ")
	fmt.Scan(&c)

	delta := b*b - 4*a*c

	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		raiz := -b / (2 * a)
		fmt.Printf("RAIZ ÚNICA: %.2f\n", raiz)
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("RAÍZES DISTINTAS: %.2f e %.2f\n", x1, x2)
	}
}
