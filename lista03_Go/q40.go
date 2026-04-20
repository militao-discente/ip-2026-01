package main

import "fmt"

func main() {
	preco := 6.0
	ingressoBase := 130
	despesa := 300.0
	maisLucro := -999999.0
	precoOtimo := 0.0
	ingressoOtimo := 0
	fmt.Println("Preco | Ingressos | Lucro")
	for preco >= 1.0 {
		ingresso := ingressoBase + int((6.0-preco)/0.6*30)
		lucro := preco*float64(ingresso) - despesa
		fmt.Printf("%.2f | %d | %.2f\n", preco, ingresso, lucro)
		if lucro > maisLucro {
			maisLucro = lucro
			precoOtimo = preco
			ingressoOtimo = ingresso
		}
		preco -= 0.6
	}
	fmt.Printf("Lucro maximo: %.2f no preco %.2f com %d ingressos\n", maisLucro, precoOtimo, ingressoOtimo)
}
