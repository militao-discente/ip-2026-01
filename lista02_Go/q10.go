package main

import "fmt"

func main() {
	var destino, iev int
	var valor float64

	fmt.Println("Escolha o destino:")
	fmt.Println("1 - região norte")
	fmt.Println("2 - região nordeste")
	fmt.Println("3 - região centro-oeste")
	fmt.Println("4 - região sul")
	fmt.Scan(&destino)

	if destino == 1 {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - somente")
		fmt.Println("2 - ida e volta")
		fmt.Scan(&iev)
		if iev == 1 {
			valor = 500.00
		} else if iev == 2 {
			valor = 900.00
		} else {
			fmt.Println("Opção inválida")
			return
		}
	} else if destino == 2 {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - somente")
		fmt.Println("2 - ida e volta")
		fmt.Scan(&iev)
		if iev == 1 {
			valor = 350.00
		} else if iev == 2 {
			valor = 650.00
		} else {
			fmt.Println("Opção inválida")
			return
		}
	} else if destino == 3 {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - somente")
		fmt.Println("2 - ida e volta")
		fmt.Scan(&iev)
		if iev == 1 {
			valor = 350.00
		} else if iev == 2 {
			valor = 600.00
		} else {
			fmt.Println("Opção inválida")
			return
		}
	} else if destino == 4 {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - somente")
		fmt.Println("2 - ida e volta")
		fmt.Scan(&iev)
		if iev == 1 {
			valor = 300.00
		} else if iev == 2 {
			valor = 550.00
		} else {
			fmt.Println("Opção inválida")
			return
		}
	} else {
		fmt.Println("Opção de destino inválida")
		return
	}

	fmt.Printf("O valor da passagem é: R$ %.2f\n", valor)
}
