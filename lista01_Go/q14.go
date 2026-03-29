package main

import (
	"fmt"
	"math"
)

func main() {
	var h, aresta float64
	fmt.Scan(&h, &aresta)
	ab := (3.0 * aresta * aresta * math.Sqrt(3.0)) / 2.0
	v := (1.0 / 3.0) * ab * h
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", v)
}
