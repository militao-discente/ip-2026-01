package main

import "fmt"

func main() {
	var b, e int
	fmt.Scan(&b, &e)

	r := 1
	for i := 0; i < e; i++ {
		r *= b
	}
	fmt.Printf("%d ", r)
}
