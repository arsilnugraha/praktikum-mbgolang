package main

import (
	"fmt"
)

func Compare(a, b string) string {
	commonSubstring := ""
	lenA := len(a)
	lenB := len(b)
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			if a[i] == b[j] {
				tempSubstring := ""
				aIndex := i
				bIndex := j
				for aIndex < lenA && bIndex < lenB && a[aIndex] == b[bIndex] {
					tempSubstring += string(a[aIndex])
					aIndex++
					bIndex++
				}
				if len(tempSubstring) > len(commonSubstring) {
					commonSubstring = tempSubstring
				}
			}
		}
	}
	return commonSubstring
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))      // Output: AKA
	fmt.Println(Compare("KANGOORO", "KANG"))   // Output: KANG
	fmt.Println(Compare("KI", "KIJANG"))       // Output: KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))  // Output: KUPU
	fmt.Println(Compare("ILALANG", "ILA"))     // Output: ILA
}