package main

import "fmt"

func playWithAsterik(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterik(5)
}