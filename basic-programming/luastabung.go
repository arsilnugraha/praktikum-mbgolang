package main

import (
	"fmt"
)

func main() {
	// input
	T := 20.0
	r := 4.0

	phi := 3.14

	luas := 2 * phi * r * (r + T)

	fmt.Printf("Luas tabung dengan tinggi %.2f dan jari-jari %.2f adalah %.2f\n", T, r, luas)
}

