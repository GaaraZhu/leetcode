package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{-2, -1}))
	fmt.Println(maxSubArray([]int{-2, -1, -3}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	var result = nums[0]
	for i := 0; i < len(nums); i++ {
		var currentTotal = nums[i]
		var lastTotal = nums[i]
		for j := i + 1; j < len(nums); j++ {
			lastTotal = lastTotal + nums[j]
			if lastTotal > currentTotal {
				currentTotal = lastTotal
			}
		}

		if currentTotal > result {
			result = currentTotal
		}
	}

	return result
}
