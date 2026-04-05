package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scan(&num)
	if num >= 0 {
		fmt.Print(math.Sqrt(num))
	} else {
		fmt.Print(num * num)
	}
}
