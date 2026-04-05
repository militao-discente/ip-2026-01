package main

import (
	"fmt"
	"time"
)

func main() {
	var matricula int
	var sm, slvhe, sb, slq, inss, ir float64
	var he string
	vhe := 10.0

	fmt.Print("Qual a matricula do funcionário: ")
	fmt.Scan(&matricula)
	fmt.Print("Digite a quantidade de horas extras do funcionário (HH:MM): ")
	fmt.Scan(&he)

	t, err := time.Parse("15:04", he)
	if err != nil {
		fmt.Println("Erro: Use o formato 00:00")
		return
	}

	hd := float64(t.Hour()) + (float64(t.Minute()) / 60.0)

	sm = 788.0
	slvhe = vhe * hd
	sb = 3.0*sm + slvhe

	if sb > 0 && sb <= 1500 {
		slq = sb
		fmt.Printf("O salário bruto do funcionário %d é: R$%.2f e o salário liquído é: R$%.2f", matricula, sb, slq)
	} else if sb > 1500.00 && sb <= 2000.00 {
		inss = sb * 0.12
		slq = sb - inss
		fmt.Printf("O salário bruto do funcionário %d é: R$%.2f e o salário liquído é: R$%.2f", matricula, sb, slq)
	} else if sb > 2000.00 {
		inss = sb * 0.12
		ir = sb * 0.20
		slq = sb - (inss + ir)
		fmt.Printf("O salário bruto do funcionário %d é: R$%.2f e o salário liquído é: R$%.2f", matricula, sb, slq)
	} else {
		fmt.Print("error")
	}

}
