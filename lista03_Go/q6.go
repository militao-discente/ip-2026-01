package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um numero inteiro positivo: ")
	fmt.Scan(&num)
	if num <= 0 {
		fmt.Println("Numero invalido")
		return
	}
	ehTriangular := false
	for a := 1; a <= num; a++ {
		for b := a + 1; b <= num; b++ {
			for c := b + 1; c <= num; c++ {
				if a*b*c == num {
					ehTriangular = true
					break
				}
			}
			if ehTriangular {
				break
			}
		}
		if ehTriangular {
			break
		}
	}
	if ehTriangular {
		fmt.Printf("%d eh triangular\n", num)
	} else {
		fmt.Printf("%d nao eh triangular\n", num)
	}
}
