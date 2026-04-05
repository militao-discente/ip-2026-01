package main

import "fmt"

func main() {
	var ide, n1, n2, n3, media, mf float64

	fmt.Print("Escreva o número de indentificação do aluno: ")
	fmt.Scan(&ide)
	fmt.Print("Escreva a primaira nota do aluno: ")
	fmt.Scan(&n1)
	fmt.Print("Escreva a segunda nota do aluno: ")
	fmt.Scan(&n2)
	fmt.Print("Escreva a terceira nota do aluno: ")
	fmt.Scan(&n3)

	media = (n1 + n2 + n3) / 3.0
	mf = (n1 + n2*2.0 + n3*3.0 + media) / 7.0

	if mf >= 9.0 && mf <= 10.0 {
		fmt.Print("O numero do aluno é: ", ide, "\n")
		fmt.Printf("As notas do aluno são: %.1f %.1f %.1f \n", n1, n2, n3)
		fmt.Printf("A media dos exercicios do aluno é: %.1f \n", media)
		fmt.Printf("A media de aproveitamento do aluno é: %.1f \n", mf)
		fmt.Print("O conceito correspondente do aluno é: A e o aluno está APROVADO!")
	} else if mf >= 7.5 && mf < 9.0 {
		fmt.Print("O numero do aluno é: ", ide, "\n")
		fmt.Printf("As notas do aluno são: %.1f %.1f %.1f \n", n1, n2, n3)
		fmt.Printf("A media dos exercicios do aluno é: %.1f \n", media)
		fmt.Printf("A media de aproveitamento do aluno é: %.1f \n", mf)
		fmt.Print("O conceito correspondente do aluno é: B e o aluno está APROVADO!")
	} else if mf >= 6.0 && mf < 7.5 {
		fmt.Print("O numero do aluno é: ", ide, "\n")
		fmt.Printf("As notas do aluno são: %.1f %.1f %.1f \n", n1, n2, n3)
		fmt.Printf("A media dos exercicios do aluno é: %.1f \n", media)
		fmt.Printf("A media de aproveitamento do aluno é: %.1f \n", mf)
		fmt.Print("O conceito correspondente do aluno é: C e o aluno está APROVADO!")
	} else if mf >= 4.0 && mf < 6.0 {
		fmt.Print("O numero do aluno é: ", ide, "\n")
		fmt.Printf("As notas do aluno são: %.1f %.1f %.1f \n", n1, n2, n3)
		fmt.Printf("A media dos exercicios do aluno é: %.1f \n", media)
		fmt.Printf("A media de aproveitamento do aluno é: %.1f \n", mf)
		fmt.Print("O conceito correspondente do aluno é: D e o aluno está REPROVADO!")
	} else if mf >= 0 && mf < 4.0 {
		fmt.Print("O numero do aluno é: ", ide, "\n")
		fmt.Printf("As notas do aluno são: %.1f %.1f %.1f \n", n1, n2, n3)
		fmt.Printf("A media dos exercicios do aluno é: %.1f \n", media)
		fmt.Printf("A media de aproveitamento do aluno é: %.1f \n", mf)
		fmt.Print("O conceito correspondente do aluno é: E e o aluno está REPROVADO!")
	} else {
		fmt.Print("Use numeros de 0 a 10 para representar as notas")
	}

}
