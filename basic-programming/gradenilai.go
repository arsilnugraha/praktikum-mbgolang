package main

import (
	"fmt"
)

func main() {
	var studentScore int = 80

	switch {
		case studentScore >= 80 && studentScore <= 100:
			fmt.Println("A")
		case studentScore >= 65 && studentScore <= 79:
			fmt.Println("B")
		case studentScore >= 50 && studentScore <= 64:
			fmt.Println("C")
		case studentScore >= 35 && studentScore <= 49:
			fmt.Println("D")
		case studentScore >= 0 && studentScore <= 34:
			fmt.Println("E")
		default:
			fmt.Println("Nilai Invalid")
	}
}