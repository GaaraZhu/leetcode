package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{1, 2, 2, 4, 5}))
	fmt.Println(findDisappearedNumbers([]int{1, 2, 2, 4, 5, 5}))
	fmt.Println(findDisappearedNumbers([]int{1, 2, 2, 4, 5, 5, 6, 6}))
}

func findDisappearedNumbers(nums []int) []int {
    if len(nums) == 0 {
		return []int{}
	}

	var res = make([]int, len(nums)+1)
	for i:=0;i<=len(nums)-1;i++ {
		res[nums[i]] = 1
	}

	var nextOverrideIndex = 0
	for i:=1;i<=len(nums);i++ {
		if res[i] == 0 {
			res[nextOverrideIndex] = i
			nextOverrideIndex = nextOverrideIndex + 1
		}
	}

	return res[0:nextOverrideIndex]
}