package main

import "fmt"

func palindrome(input string) bool {
	// Menghilangkan spasi dan mengubah semua huruf menjadi lowercase
	input = removeSpacesAndToLower(input)

	// Membandingkan karakter di awal dan akhir string
	// dengan iterasi dari kedua arah
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}

	return true
}

func removeSpacesAndToLower(input string) string {
	output := ""
	for i := 0; i < len(input); i++ {
		if input[i] != ' ' {
			output += string(input[i])
		}
	}
	return output
}

func main() {
	fmt.Println(palindrome("civic"))          // true
	fmt.Println(palindrome("katak"))          // true
	fmt.Println(palindrome("kasur rusak"))    // true
	fmt.Println(palindrome("mistar"))         // false
	fmt.Println(palindrome("lion"))           // false
}
