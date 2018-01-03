package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityElement([]int{2,2,3,3,2,2,2,6}))
}

func majorityElement(nums []int) int {
	cm := map[int]int{}
	for _, n := range nums {
		c, ok := cm[n]
		if !ok {
			cm[n] = 1
		} else {
			cm[n] = c + 1
		}
	}
    
	for k, v := range cm {
		if v > len(nums) /2 {
			return k
		}
	}

	return -1
}

