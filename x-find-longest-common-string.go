package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonString("abcdcg", "esighcdcjsl"))
}

func longestCommonString(s1, s2 string) string {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	var resultS = ""
	var currentS = ""
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if runes1[i] == runes2[j] {
				currentS = string(runes1[i])
				for k := 1; (i+k < len(s1)) && (j+k < len(s2)); k++ {
					if runes1[i+k] != runes2[j+k] {
						break
					}
					currentS += string(runes1[i+k])
				}

				if len(currentS) > len(resultS) {
					resultS = currentS
				}
			}
			currentS = ""
		}
	}

	return resultS
}
