package main

import "fmt"

func main() {
	var vlr, vf float64
	var pag int
	fmt.Print("Digite o preço normal da etiqueta: ")
	fmt.Scan(&vlr)
	fmt.Print("Escolha a forma de pagamento: 1- dinheiro ou cheque 2- crédito 3- Em duas vzes 4- Em trêS vezes:  ")
	fmt.Scan(&pag)
	if pag == 1 {
		vf = vlr - (vlr * 0.10)
		fmt.Printf("O valor a ser pago pelo produto considerando o preço normal de etiqueta e a escolha da condição de pagamento é: R$%.2f", vf)
	} else if pag == 2 {
		vf = vlr - (vlr * 0.05)
		fmt.Printf("O valor a ser pago pelo produto considerando o preço normal de etiqueta e a escolha da condição de pagamento é: R$%.2f", vf)
	} else if pag == 3 {
		vf = vlr
		fmt.Printf("O valor a ser pago pelo produto considerando o preço normal de etiqueta e a escolha da condição de pagamento é: R$%.2f", vf)
	} else if pag == 4 {
		vf = vlr + (vlr * 0.10)
		fmt.Printf("O valor a ser pago pelo produto considerando o preço normal de etiqueta e a escolha da condição de pagamento é: R$%.2f", vf)
	} else {
		fmt.Print("Opção inválida!")
	}
}
