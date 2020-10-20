package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(longestPalindrome("aac"))
	// fmt.Println(longestPalindrome("abc"))
	// fmt.Println(longestPalindrome("abcbdbcgg"))
	// fmt.Println(longestPalindrome("aacdefcaa"))
	fmt.Println(longestPalindrome("7aba"))
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

	return strings.Replace(resultS, "#", "", -1)
}

// Time Limit Exceeded
func longestPalindrome2(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	runes1 := []rune(s)
	var runes2 []rune
	for i := 0; i < length; i++ {
		runes2 = append(runes2, runes1[length-i-1])
	}

	var resultS = ""
	var currentS = ""
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if runes1[i] == runes2[j] {
				currentS = string(runes1[i])
				var endIndex int
				for k := 1; (i+k < length) && (j+k < length); k++ {
					if runes1[i+k] != runes2[j+k] {
						break
					}
					endIndex = j + k
					currentS += string(runes1[i+k])
				}
				if len(currentS) > len(resultS) && endIndex+i == length-1 {
					resultS = currentS
				}
			}
			currentS = ""
		}
	}

	return resultS
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
