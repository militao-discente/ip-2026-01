package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	media := (a + b + c) / 3.0
	fmt.Printf("MEDIA = %.2f\n", media)
	if media >= 6.0 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
