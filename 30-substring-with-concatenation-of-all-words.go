package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("123"[0, 1])
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) < len(words)*len(words[0]) {
		return []int{}
	}

	var concatenation = []string{}
	for i := 0; i < len(words); i++ {
		if i == 0 {
			concatenation = append(concatenation, words[0])
			continue
		}

		var newConcatenations = []string{}
		for _, c := range concatenation {
			newConcatenations = append(newConcatenations, c+words[i])
			newConcatenations = append(newConcatenations, words[i]+c)
			// for j:=0;j<len(c);{
			// 	c[0, j] + words[i] + c[j+1]
			// 	j = j + len(words[i])
			// }
		}

		concatenation = newConcatenations
	}

	var indexMap = make(map[int]bool)
	for _, c := range concatenation {
		fmt.Println(c)
		if i := strings.Index(s, c); i != -1 {
			indexMap[i] = true
		}
	}

	var indices = make([]int, 0, len(indexMap))
	for k := range indexMap {
		indices = append(indices, k)
	}

	return indices
}
