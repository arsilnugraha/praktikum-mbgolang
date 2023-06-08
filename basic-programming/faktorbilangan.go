package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanf("%d", &bilangan)

	fmt.Println("Faktor dari", bilangan, "adalah:")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}
