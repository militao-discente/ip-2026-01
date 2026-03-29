package main

import "fmt"

func main() {
	var salario, kw float64
	fmt.Scan(&salario, &kw)
	custoKw := (salario * 0.7) / 100.0
	consumo := custoKw * kw
	comDesconto := consumo * 0.9
	fmt.Printf("Custo por kW: R$ %.2f\n", custoKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", consumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", comDesconto)
}
