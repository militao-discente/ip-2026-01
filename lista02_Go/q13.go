package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro de 3 dígitos: ")
	fmt.Scan(&num)
	if num >= 100 && num <= 999 {
		segundoDigito := (num / 10) % 10
		fmt.Printf("O segundo dígito do número %d é: %d\n", num, segundoDigito)
	} else {
		fmt.Println("Por favor, digite um número inteiro de 3 dígitos.")
	}
}
