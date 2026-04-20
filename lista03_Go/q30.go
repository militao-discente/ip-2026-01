package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Raio (cm) | Volume (cm3)")
	for r := 0.0; r <= 20.0; r += 0.5 {
		vol := (4.0 / 3.0) * math.Pi * r * r * r
		fmt.Printf("%.1f | %.2f\n", r, vol)
	}
}
