package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)
	if 0 <= idade && idade <= 2 {
		fmt.Println("Recém nascido.")
	} else if 3 <= idade && idade <= 11 {
		fmt.Println("Criança.")
	} else if 12 <= idade && idade <= 19 {
		fmt.Println("Adolescente.")
	} else if 20 <= idade && idade <= 55 {
		fmt.Println("Adulto.")
	} else if idade > 55 {
		fmt.Println("Idoso.")
	}
}
