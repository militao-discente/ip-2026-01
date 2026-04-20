package main

import "fmt"

func main() {
	var decimal int
	fmt.Print("Digite numero decimal: ")
	fmt.Scan(&decimal)
	if decimal == 0 {
		fmt.Println("0")
		return
	}
	binario := ""
	for decimal > 0 {
		binario = fmt.Sprintf("%d", decimal%2) + binario
		decimal /= 2
	}
	fmt.Printf("Binario: %s\n", binario)
}
