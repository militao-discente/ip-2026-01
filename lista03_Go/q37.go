package main

import "fmt"

func main() {
	var octalStr string
	fmt.Print("Digite numero na base 8: ")
	fmt.Scan(&octalStr)
	decimal := 0
	for _, digit := range octalStr {
		decimal = decimal*8 + int(digit-'0')
	}
	fmt.Printf("Decimal: %d\n", decimal)
}
