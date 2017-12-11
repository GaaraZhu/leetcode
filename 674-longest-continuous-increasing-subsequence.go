package main

import "fmt"

func main() {
	// fmt.Println(findLengthOfLCIS([]int{}))
	fmt.Println(findLengthOfLCIS([]int{1}))
	// fmt.Println(findLengthOfLCIS([]int{0, 0}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))
	// fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	// fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
	// fmt.Println(findLengthOfLCIS([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	// fmt.Println(findLengthOfLCIS([]int{2147483646, 2147483644, 2147483645, 0}))
	// fmt.Println(findLengthOfLCIS([]int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645}))
}

func findLengthOfLCIS(nums []int) (max int) {
	var currentMax int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			currentMax = 1
		}

		if i-1 >= 0 && nums[i-1] < nums[i] {
			currentMax++
		}

		if (i-1 >= 0 && nums[i-1] >= nums[i]) || i == len(nums)-1 {
			if currentMax > max {
				max = currentMax
			}
			currentMax = 1
		}

	}
	return max
}
