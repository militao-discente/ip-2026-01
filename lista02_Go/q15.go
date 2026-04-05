package main

import "fmt"

func main() {
	dia, mes, ano := 0, 0, 0
	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)
	fmt.Print("Digite o mês: ")
	fmt.Scan(&mes)
	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

	meses := []string{
		"",
		"janeiro",
		"fevereiro",
		"março",
		"abril",
		"maio",
		"junho",
		"julho",
		"agosto",
		"setembro",
		"outubro",
		"novembro",
		"dezembro",
	}

	if mes >= 1 && mes <= 12 {
		fmt.Printf("Data informada: %02d de %s de %d\n", dia, meses[mes], ano)
	} else {
		fmt.Println("Mês inválido!")
	}
}
