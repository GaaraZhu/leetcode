package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{}))
	fmt.Println(findMin([]int{1,2,3,0}))
	fmt.Println(findMin([]int{4,5,6,1,2,3}))
	fmt.Println(findMin([]int{4,5,6,2,3}))
	fmt.Println(findMin([]int{0, 1, -1}))
}

func findMin(nums []int) int {
	var min = 0 
	for i:=0 ; i < len(nums); i++ {
		if i == 0 {
			min = nums[0]
			continue
		}

		if nums[i] < min {
			min = nums[i]
			break
		}
	}
	return min
}
