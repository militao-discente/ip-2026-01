package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if 20 < num && num < 90 {
		fmt.Println("O número está entre 20 e 90.")
	} else {
		fmt.Println("O número não está entre 20 e 90.")
	}
}
