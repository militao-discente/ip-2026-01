package main

import "fmt"

func main() {
	var tipo string
	var taxa, metros, conta float64

	fmt.Print("Escreva o tipo do consumidor - residencial, comercial ou industrial:\n")
	fmt.Scan(&tipo)
	if tipo == "residencial" {
		taxa = 5.00
		fmt.Print("Quantos metros deseja caucular?\n")
		fmt.Scan(&metros)
		if metros > 0 {
			conta = taxa + metros*0.05
			fmt.Printf("%.2f + %.2f * 0.05 = R$ %.2f", taxa, metros, conta)
		} else {
			fmt.Print("Digite um numero positivo")
		}
	} else if tipo == "comercial" {
		taxa = 500.00
		fmt.Print("Quantos metros deseja caucular?\n")
		fmt.Scan(&metros)
		if metros > 0 && metros < 80 {
			conta = taxa + metros*0.0
			fmt.Printf("%.2f + %.2f * 0.0 = R$ %.2f", taxa, metros, conta)
		} else if metros > 80 {
			conta = taxa + metros*0.25
			fmt.Printf("%.2f + %.2f * 0.25 = R$ %.2f", taxa, metros, conta)
		} else {
			fmt.Print("Digite um numero positivo")
		}
	} else if tipo == "industrial" {
		taxa = 800.00
		fmt.Print("Quantos metros deseja caucular?\n")
		fmt.Scan(&metros)
		if metros > 0 && metros < 100 {
			conta = taxa + metros*0.0
			fmt.Printf("%.2f + %.2f * 0.0 = R$ %.2f", taxa, metros, conta)
		} else if metros > 100 {
			conta = taxa + metros*0.04
			fmt.Printf("%.2f + %.2f * 0.04 = R$ %.2f", taxa, metros, conta)
		} else {
			fmt.Print("Digite um numero positivo")
		}
	} else {
		fmt.Print("Resposta invalida")
	}
}
