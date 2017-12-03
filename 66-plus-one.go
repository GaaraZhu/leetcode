package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{1, 4, 5, 7}))
}

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	digits[len(digits)-1] = digits[len(digits)-1] + 1
    for i:=len(digits)-1; i>=0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				return append([]int{1}, digits...)
			}
			digits[i-1] =  digits[i-1] + 1
		} else {
			return digits
		}
	}

	return digits
}
