package main

import "fmt"

func main() {
	var A, B, C, MENOR, INTER, MAIOR int
	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&A)
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&B)
	if A == B {
		fmt.Printf("ERRO: Os números devem ser diferentes\n")
	} else {
		fmt.Print("Digite o valor de C: ")
		fmt.Scan(&C)
		if A == B || A == C || B == C {
			fmt.Printf("ERRO: Os números devem ser diferentes\n")
		} else {
			if A < B && A < C {
				MENOR = A
				if B < C {
					INTER = B
					MAIOR = C
				} else {
					INTER = C
					MAIOR = B
				}
			} else if B < A && B < C {
				MENOR = B
				if A < C {
					INTER = A
					MAIOR = C
				} else {
					INTER = C
					MAIOR = A
				}
			} else {
				MENOR = C
				if A < B {
					INTER = A
					MAIOR = B
				} else {
					INTER = B
					MAIOR = A
				}
			}
			fmt.Printf("%d, %d, %d\n", MENOR, INTER, MAIOR)
		}

	}
}
