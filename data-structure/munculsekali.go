package main

import (
	"fmt"
)

func munculSekali(input string) []int {
	countMap := make(map[rune]int)

	for _, digit := range input {
		countMap[digit]++
	}

	uniqueNumbers := make([]int, 0, len(countMap))

	for digit, count := range countMap {
		if count == 1 {
			number := int(digit - '0')
			uniqueNumbers = append(uniqueNumbers, number)
		}
	}

	return uniqueNumbers
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}