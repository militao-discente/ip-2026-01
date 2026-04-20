package main

import "fmt"

func main() {
	soma := 0
	qtde := 0
	somaPares := 0
	qtdePares := 0
	qtdeImpares := 0
	maior := -999999
	menor := 999999
	var num int
	for {
		fmt.Print("Digite um numero (30000 para terminar): ")
		fmt.Scan(&num)
		if num == 30000 {
			break
		}
		soma += num
		qtde++
		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}
		if num%2 == 0 {
			somaPares += num
			qtdePares++
		} else {
			qtdeImpares++
		}
	}
	fmt.Printf("Soma: %d\n", soma)
	fmt.Printf("Quantidade: %d\n", qtde)
	if qtde > 0 {
		media := float64(soma) / float64(qtde)
		fmt.Printf("Media: %.2f\n", media)
	} else {
		fmt.Println("Media: 0.00")
	}
	fmt.Printf("Maior: %d\n", maior)
	fmt.Printf("Menor: %d\n", menor)
	if qtdePares > 0 {
		mediaPares := float64(somaPares) / float64(qtdePares)
		fmt.Printf("Media pares: %.2f\n", mediaPares)
	} else {
		fmt.Println("Media pares: 0.00")
	}
	porcImpares := 0.0
	if qtde > 0 {
		porcImpares = float64(qtdeImpares) * 100.0 / float64(qtde)
	}
	fmt.Printf("Porcentagem impares: %.2f%%\n", porcImpares)
}
