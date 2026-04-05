package main

import "fmt"

func main() {
	var num1, num2, soma int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	soma = num1 + num2
	if soma > 20 {
		fmt.Println(soma + 8)
	} else {
		fmt.Println(soma - 5)
	}
}
