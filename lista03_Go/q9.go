package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite N: ")
	fmt.Scan(&n)
	aprovados := 0
	exames := 0
	reprovados := 0
	somaMedias := 0.0
	for i := 0; i < n; i++ {
		var nota1, nota2 float64
		fmt.Printf("Aluno %d - Nota 1: ", i+1)
		fmt.Scan(&nota1)
		fmt.Printf("Aluno %d - Nota 2: ", i+1)
		fmt.Scan(&nota2)
		media := (nota1 + nota2) / 2.0
		somaMedias += media
		fmt.Printf("Media: %.2f - ", media)
		if media <= 3 {
			fmt.Println("Reprovado")
			reprovados++
		} else if media < 7 {
			fmt.Println("Exame")
			exames++
		} else {
			fmt.Println("Aprovado")
			aprovados++
		}
	}
	fmt.Printf("Aprovados: %d\n", aprovados)
	fmt.Printf("Exame: %d\n", exames)
	fmt.Printf("Reprovados: %d\n", reprovados)
	if n > 0 {
		mediaClasse := somaMedias / float64(n)
		fmt.Printf("Media da classe: %.2f\n", mediaClasse)
	}
}
