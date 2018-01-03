package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPeakElement([]int{}))
	fmt.Println(findPeakElement([]int{1}))
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 3, 1, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 3, 1, 2, 3, 1, 4, 0}))
}

func findPeakElement(nums []int) int {
	var peak  int
	var peakIndex = -1
	for i:=0; i<len(nums);i++ {
		if i==0 {
			peak = nums[0]
			peakIndex = 0
			continue
		}

		if nums[i] > nums[i-1] {
			if i+1 < len(nums) {
				if nums[i+1] > nums[i] {
					continue
				}
			}
			if nums[i] > peak {
				peak = nums[i]
				peakIndex = i
			}
		}
	}
	
	return peakIndex
}

