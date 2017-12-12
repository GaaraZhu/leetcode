package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))                                                                                                                   //["eat","oath"]
	fmt.Println(findWords([][]byte{{'a'}}, []string{"a"}))                                                                                                 //["a"]
	fmt.Println(findWords([][]byte{{'a', 'a'}}, []string{"aaa"}))                                                                                          //[]
	fmt.Println(findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"acdb"}))                                                                             //["acdb"]
	fmt.Println(findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"ab", "cb", "ad", "bd", "ac", "ca", "da", "bc", "db", "adcb", "dabc", "abb", "acb"})) //["ab","ac","bd","ca","db"]
	fmt.Println(findWords([][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"ab"}))                                                                               //["ab"]
}

func findWords(board [][]byte, words []string) []string {
	dict := map[string]int{}
	for _, w := range words {
		if w == "" {
			continue
		}
		for i := 1; i <= len(w); i++ {
			v, ok := dict[w[0:i]]
			if !ok || v != 1 {
				dict[w[0:i]] = 0
			}
		}
		dict[w] = 1
	}
	var res = map[string]interface{}{}
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			res = checkLetter("", j, i, [][]int{}, dict, board, res)
		}
	}

	var resArry []string
	for s := range res {
		resArry = append(resArry, s)
	}
	return resArry
}

func checkLetter(existing string, j, i int, usedCells [][]int, dict map[string]int, board [][]byte, result map[string]interface{}) map[string]interface{} {
	if i < 0 || i >= len(board[0]) || j < 0 || j >= len(board) || isCellUsed([]int{j, i}, usedCells) {
		return result
	}
	newW := existing + string(board[j][i])
	v, ok := dict[newW]
	if !ok {
		return result
	}

	if v == 1 {
		result[newW] = struct{}{}
	}
	usedCells = append(usedCells, []int{j, i})
	result = checkLetter(newW, j, i-1, usedCells, dict, board, result)
	result = checkLetter(newW, j, i+1, usedCells, dict, board, result)
	result = checkLetter(newW, j-1, i, usedCells, dict, board, result)
	return checkLetter(newW, j+1, i, usedCells, dict, board, result)
}

func isCellUsed(cell []int, usedCells [][]int) bool {
	for _, c := range usedCells {
		if c[0] == cell[0] && c[1] == cell[1] {
			return true
		}
	}

	return false
}
