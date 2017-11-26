package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits(5))
}

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	var res = []int{0, 1}
	for i:=2; i<= num; i++ {
		if i%2 == 0 {
			res = append(res, res[i/2])
		} else {
			res = append(res, res[i/2]+1)
		}
	}
	return res
}

