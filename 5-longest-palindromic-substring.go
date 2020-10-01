package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestPalindromeImproved("aac"))
	fmt.Println(longestPalindromeImproved("abc"))
	fmt.Println(longestPalindromeImproved("abcbdbcgg"))
	fmt.Println(longestPalindromeImproved("aacdefcaa"))
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

func longestPalindromeImproved(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	var tmpBytes []byte
	for i := 0; i < length; i++ {
		tmpBytes = append(tmpBytes, '#')
		tmpBytes = append(tmpBytes, s[i])
	}
	tmpBytes = append(tmpBytes, '#')

	// find the longgest common string
	var resultS = ""
	var currentS = ""
	for i := 0; i < len(tmpBytes); i++ {
		currentS = string(tmpBytes[i])
		for j := 1; (i-j > 0) && (i+j < len(tmpBytes)); j++ {
			if tmpBytes[i-j] != tmpBytes[i+j] {
				break
			}
			currentS = string(tmpBytes[i-j]) + currentS + string(tmpBytes[i+j])
		}
		if len(currentS) > len(resultS) {
			resultS = currentS
		}
		currentS = ""
	}

	return strings.Replace(resultS, '#')
}
