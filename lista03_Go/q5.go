package main

import "fmt"

func main() {
	qtdeMais50 := 0
	somaAlturas := 0.0
	qtdeAlturas := 0
	qtdePesoMenor40 := 0
	qtdeTotal := 0
	for {
		var idade int
		var altura, peso float64
		fmt.Print("Digite idade: ")
		fmt.Scan(&idade)
		fmt.Print("Digite altura: ")
		fmt.Scan(&altura)
		fmt.Print("Digite peso: ")
		fmt.Scan(&peso)
		qtdeTotal++
		if idade > 50 {
			qtdeMais50++
		}
		if idade >= 10 && idade <= 20 {
			somaAlturas += altura
			qtdeAlturas++
		}
		if peso < 40 {
			qtdePesoMenor40++
		}
		var continuar int
		fmt.Print("Deseja continuar (1-Sim, outro-Nao)? ")
		fmt.Scan(&continuar)
		if continuar != 1 {
			break
		}
	}
	fmt.Printf("Pessoas com mais de 50 anos: %d\n", qtdeMais50)
	if qtdeAlturas > 0 {
		mediaAlturas := somaAlturas / float64(qtdeAlturas)
		fmt.Printf("Media de alturas (10-20 anos): %.2f\n", mediaAlturas)
	} else {
		fmt.Println("Media de alturas (10-20 anos): 0.00")
	}
	porcentagem := 0.0
	if qtdeTotal > 0 {
		porcentagem = float64(qtdePesoMenor40) * 100.0 / float64(qtdeTotal)
	}
	fmt.Printf("Porcentagem peso < 40kg: %.2f%%\n", porcentagem)
}
