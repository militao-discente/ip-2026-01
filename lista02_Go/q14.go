package main

import "fmt"

func main() {
	var vlr float64
	var opcao string
	fmt.Print("Digite o valor inicial do carro: ")
	fmt.Scan(&vlr)
	fmt.Print(`O carro pode ter as seguintes opções:
	 a) Ar condicionado (R$ 1750,00)
	 b) Pintura metálica (R$ 800,00)
	 c) Vidro elétrico (R$ 1200,00)
	 d) Direção hidráulica (R$ 2000,00)
	 Digite a letra para adicionar a opção escolhida, caso não queira nenhuma digite x: `)
	fmt.Scan(&opcao)
	if opcao == "a" {
		vlr += 1750.00
		fmt.Printf("O valor do carro com a opção escolhida é: %.2f\n", vlr)
	} else if opcao == "b" {
		vlr += 800.00
		fmt.Printf("O valor do carro com a opção escolhida é: %.2f\n", vlr)
	} else if opcao == "c" {
		vlr += 1200.00
		fmt.Printf("O valor do carro com a opção escolhida é: %.2f\n", vlr)
	} else if opcao == "d" {
		vlr += 2000.00
		fmt.Printf("O valor do carro com a opção escolhida é: %.2f\n", vlr)
	} else if opcao == "x" {
		fmt.Printf("O valor do carro sem opções é: %.2f\n", vlr)
	} else {
		fmt.Println("Opção inválida.")
	}
}
