package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("aac"))
	fmt.Println(longestPalindrome("abc"))
	fmt.Println(longestPalindrome("abcbdbcgg"))
	fmt.Println(longestPalindrome("aacdefcaa"))
}

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	runes := []rune(s)
	var tmpRunes []rune
	for i := 0; i < length; i++ {
		tmpRunes = append(tmpRunes, '#')
		tmpRunes = append(tmpRunes, runes[i])
	}
	tmpRunes = append(tmpRunes, '#')

	// find the longgest common string
	var resultS = ""
	var currentS = ""
	for i := 0; i < len(tmpRunes); i++ {
		currentS = string(tmpRunes[i])
		for j := 1; (i-j > 0) && (i+j < len(tmpRunes)); j++ {
			if tmpRunes[i-j] != tmpRunes[i+j] {
				break
			}
			currentS = string(tmpRunes[i-j]) + currentS + string(tmpRunes[i+j])
		}
		if len(currentS) > len(resultS) {
			resultS = currentS
		}
		currentS = ""
	}

	rRunes := []rune(resultS)
	var rsRunes []rune
	for i := 0; i < len(rRunes); i++ {
		if rRunes[i] != '#' {
			rsRunes = append(rsRunes, rRunes[i])
		}
	}

	return string(rsRunes)
}
