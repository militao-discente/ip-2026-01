package main

import "fmt"

func main() {
	var cpf [11]int
	fmt.Print("Digite os 11 digitos do CPF: ")
	for i := 0; i < 11; i++ {
		fmt.Scan(&cpf[i])
	}
	soma1 := 0
	for i := 0; i < 9; i++ {
		soma1 += cpf[i] * (10 - i)
	}
	resto1 := soma1 % 11
	dv1 := 0
	if resto1 >= 2 {
		dv1 = 11 - resto1
	}
	soma2 := 0
	for i := 0; i < 9; i++ {
		soma2 += cpf[i] * (11 - i)
	}
	soma2 += dv1 * 2
	resto2 := soma2 % 11
	dv2 := 0
	if resto2 >= 2 {
		dv2 = 11 - resto2
	}
	if dv1 == cpf[9] && dv2 == cpf[10] {
		fmt.Println("CPF valido")
	} else {
		fmt.Println("CPF invalido")
	}
}
