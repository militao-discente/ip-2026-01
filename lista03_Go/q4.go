package main

import "fmt"

func main() {
	var num int
	for {
		fmt.Print("Digite um numero inteiro (0 ou negativo para terminar): ")
		fmt.Scan(&num)
		if num <= 0 {
			break
		}
		ehQuadrado := false
		for i := 1; i*i <= num; i++ {
			if i*i == num {
				ehQuadrado = true
				break
			}
		}
		if ehQuadrado {
			fmt.Printf("%d eh quadrado perfeito\n", num)
		} else {
			fmt.Printf("%d nao eh quadrado perfeito\n", num)
		}
	}
}
