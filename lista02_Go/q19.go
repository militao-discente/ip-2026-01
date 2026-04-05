package main

import (
	"fmt"
	"math"
)

func main() {
	var escolha1, escolha2 int
	var raio, altura float64

	fmt.Print("Escoha a opção que deseja calcular 1-cone 2-cilindro 3-esfera: ")
	fmt.Scan(&escolha1)

	if escolha1 == 1 {
		fmt.Print("Escolha oque deseja calcular 1-volume 2-área: ")
		fmt.Scan(&escolha2)
		if escolha2 == 1 {
			fmt.Print("Digite o raio do cone: ")
			fmt.Scan(&raio)
			fmt.Print("Digite a altura do cone: ")
			fmt.Scan(&altura)
			volume := (math.Pi * (raio * raio) * altura) / 3.0
			fmt.Print(volume)
		} else if escolha2 == 2 {
			fmt.Print("Digite o raio do cone: ")
			fmt.Scan(&raio)
			fmt.Print("Digite a altura do cone: ")
			fmt.Scan(&altura)
			area := math.Pi * raio * math.Sqrt((raio*raio)+(altura*altura))
			fmt.Print(area)
		} else {
			fmt.Print("Escolha inválida")
		}
	} else if escolha1 == 2 {
		fmt.Print("Escolha oque deseja calcular 1-volume 2-área: ")
		fmt.Scan(&escolha2)
		if escolha2 == 1 {
			fmt.Print("Digite o raio do cilindro: ")
			fmt.Scan(&raio)
			fmt.Print("Digite a altura do cilindro: ")
			fmt.Scan(&altura)
			volume := math.Pi * (raio * raio) * altura
			fmt.Print(volume)
		} else if escolha2 == 2 {
			fmt.Print("Digite o raio do cilindro: ")
			fmt.Scan(&raio)
			fmt.Print("Digite a altura do cilindro: ")
			fmt.Scan(&altura)
			area := 2.0 * math.Pi * raio * altura
			fmt.Print(area)
		} else {
			fmt.Print("Escolha inválida")
		}
	} else if escolha1 == 3 {
		fmt.Print("Escolha oque deseja calcular 1-volume 2-área: ")
		fmt.Scan(&escolha2)
		if escolha2 == 1 {
			fmt.Print("Digite o raio da esfera: ")
			fmt.Scan(&raio)
			volume := (4.0 / 3.0) * math.Pi * (raio * raio * raio)
			fmt.Print(volume)
		} else if escolha2 == 2 {
			fmt.Print("Digite o raio da esfera: ")
			fmt.Scan(&raio)
			area := 4.0 * math.Pi * (raio * raio)
			fmt.Print(area)
		} else {
			fmt.Print("Escolha inválida")
		}
	} else {
		fmt.Print("Escolha inválida")
	}
}
