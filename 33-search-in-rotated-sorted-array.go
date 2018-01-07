package main

import (
    "fmt"
)

func main() {
	fmt.Println(search([]int{3, 1}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 7))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 8))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 6))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, -1))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if target < nums[0] {
		for i:=0; i<len(nums); i++ {
			if nums[i] >= nums[0] {
				continue
			}

			if nums[i] == target {
				return i
			} else if nums[i] > target {
				return -1
			}
		}
	} else {
		for i:=0; i<len(nums); i++ {
			if nums[i] == target {
				return i
			}

			if nums[i]<nums[0] {
				return -1
			}
		}
	}

	return -1
}
