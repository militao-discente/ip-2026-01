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
	hex := ""
	for decimal > 0 {
		resto := decimal % 16
		var digito string
		if resto < 10 {
			digito = fmt.Sprintf("%d", resto)
		} else {
			digito = string('A' + resto - 10)
		}
		hex = digito + hex
		decimal /= 16
	}
	fmt.Printf("Hexadecimal: %s\n", hex)
}
