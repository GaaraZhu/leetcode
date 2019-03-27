package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeElement([]int{-2, -2, -1, -2, -2, -3}, -2))
}

func removeElement(nums []int, val int) int {
	j := 0
	size := len(nums)
	for i := 0; i < size && j < size; i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
