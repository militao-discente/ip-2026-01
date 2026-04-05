package main

import "fmt"

func main() {
	var valor float64
	fmt.Print("Digite o valor de compra: ")
	fmt.Scanf("%f", &valor)
	if valor < 10 {
		fmt.Printf("O valor da venda é: R$ %.2f\n", valor*1.70)
	} else if valor >= 10 && valor < 30 {
		fmt.Printf("O valor da venda é: R$ %.2f\n", valor*1.50)
	} else if valor >= 30 && valor < 50 {
		fmt.Printf("O valor da venda é: R$ %.2f\n", valor*1.40)
	} else if valor >= 50 {
		fmt.Printf("O valor da venda é: R$ %.2f\n", valor*1.30)
	} else {
		fmt.Println("Valor inválido")
	}
}
