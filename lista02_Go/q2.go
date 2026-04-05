package main

import "fmt"

func main() {
	var num float32
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println("o número é positivo")
	} else {
		fmt.Println("o número é negativo")
	}
}
