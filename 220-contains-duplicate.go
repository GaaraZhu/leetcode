package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1}, 0, 0))             //false
	fmt.Println(containsNearbyAlmostDuplicate([]int{1}, 1, 1))             //false
	fmt.Println(containsNearbyAlmostDuplicate([]int{-1, -1}, 1, 0))        // true
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 3, 8, 4, 5}, 2, 4)) //true
	fmt.Println(containsNearbyAlmostDuplicate([]int{-3, 3}, 2, 4))         //false
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 || k < 0 || len(nums) < 2 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t) {
				return true
			}
		}
	}

	return false
}
