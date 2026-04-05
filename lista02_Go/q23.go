package main

import "fmt"

func main() {
	var idd int
	fmt.Print("Qual sua idade? ")
	fmt.Scan(&idd)
	if idd > 0 && idd < 16 {
		fmt.Print("Sua classe eleitoral é: Não-eleitor")
	} else if idd >= 18 && idd <= 65 {
		fmt.Print("Sua classe eleitoral é: Eleitor obrigatório")
	} else if idd >= 16 && idd <= 18 || idd > 65 {
		fmt.Print("Sua classe eleitoral é: Eleitor facultativo")
	} else {
		fmt.Print("Erro idade inválida!")
	}
}
