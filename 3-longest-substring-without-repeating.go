package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("bbbaaa"))
}

func lengthOfLongestSubstring(s string) int {
	var left, right int
	var maxLength int
	for left, right = 0, 0; right < len(s); right++ {
		if right == 0 {
			maxLength = 1
			continue
		}

		currentSubstring := s[left:right]
		nextCharString := s[right : right+1]
		existingCharIndex := strings.Index(currentSubstring, nextCharString)
		if existingCharIndex < 0 {
			if (right - left + 1) > maxLength {
				maxLength = right - left + 1
				continue
			}
		}

		left += existingCharIndex + 1
	}

	return maxLength
}
