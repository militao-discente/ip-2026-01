package main

import "fmt"

func main() {
	maisGordoId := 0
	maisGordoPeso := 0.0
	maisMagroId := 0
	maisMagroPeso := 999999.0
	for i := 1; i <= 90; i++ {
		var id int
		var peso float64
		fmt.Printf("Boi %d - ID: ", i)
		fmt.Scan(&id)
		fmt.Printf("Boi %d - Peso: ", i)
		fmt.Scan(&peso)
		if peso > maisGordoPeso {
			maisGordoId = id
			maisGordoPeso = peso
		}
		if peso < maisMagroPeso {
			maisMagroId = id
			maisMagroPeso = peso
		}
	}
	fmt.Printf("Mais gordo - ID: %d Peso: %.2f\n", maisGordoId, maisGordoPeso)
	fmt.Printf("Mais magro - ID: %d Peso: %.2f\n", maisMagroId, maisMagroPeso)
}
