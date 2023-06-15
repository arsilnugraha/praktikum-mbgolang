package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPrime(1000000007)) // true
	fmt.Println(isPrime(13))         // true
	fmt.Println(isPrime(17))         // true
	fmt.Println(isPrime(20))         // false
	fmt.Println(isPrime(35))         // false
}