package main

import "fmt"

func ArrayMerge(arrayA []string, arrayB []string) []string {
	uniqueNames := make(map[string]bool)

	for _, name := range arrayA {
		uniqueNames[name] = true
	}

	for _, name := range arrayB {
		if !uniqueNames[name] {
			arrayA = append(arrayA, name)
			uniqueNames[name] = true
		}
	}

	return arrayA
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
