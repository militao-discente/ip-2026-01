package main

import "fmt"

func main() {
	var dia int
	var dvd string
	var preco, tipo, vlf float64

	fmt.Print("Qual preço normal do DVD ? ")
	fmt.Scan(&preco)
	if preco > 0 {
		fmt.Print("O DVD é comun ou lançamento? ")
		fmt.Scan(&dvd)
		if dvd == "comun" {
			tipo = preco
			fmt.Print("Represente o dia da semana com números de 1 á 7 conforme a sequência de dias da semana: ")
			fmt.Scan(&dia)
			if dia == 4 || dia == 6 || dia == 7 || dia == 1 {
				vlf = tipo
			} else if dia == 2 || dia == 3 || dia == 5 {
				vlf = tipo - (preco * 0.40)
			} else {
				fmt.Print("Escolha um numero de 1 á 7: ")
			}
			fmt.Printf("R$%.2f", vlf)
		} else if dvd == "lançamento" {
			tipo = preco + (preco * 0.15)
			fmt.Print(tipo)
			fmt.Print("Represente o dia da semana com números de 1 á 7 conforme a sequência de dias da semana: ")
			fmt.Scan(&dia)
			if dia == 4 || dia == 6 || dia == 7 || dia == 1 {
				vlf = tipo
			} else if dia == 2 || dia == 3 || dia == 5 {
				vlf = tipo - (preco * 0.40)
			} else {
				fmt.Print("Escolha um numero de 1 á 7: ")
			}
			fmt.Printf("R$%.2f", vlf)
		} else {
			fmt.Print("Opção inválida!")
		}
	} else {
		fmt.Print("Digite um valor inteiro e positivo")
	}
}
