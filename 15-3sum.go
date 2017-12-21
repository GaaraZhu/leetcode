package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{}))                      //[]
	fmt.Println(threeSum([]int{1, 2}))                  //[]
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))   //[[-1 0 1] [-1 -1 2]]
	fmt.Println(threeSum([]int{0, 0, 0, 1, 1, -1, -2})) //[[0 0 0] [-1 0 1] [-2 1 1]]
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	tripletsMap := map[triplets]struct{}{}
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		var j = i + 1
		var k = len(nums) - 1

		for j < k {
			switch {
			case nums[j]+nums[k] == target:
				tripletsMap[newTriplets(nums[i], nums[j], nums[k])] = struct{}{}
				j++
			case nums[j]+nums[k] > target:
				k--
			default:
				j++
			}
		}
	}

	var res = [][]int{}
	for k := range tripletsMap {
		res = append(res, []int{k.small, k.middle, k.large})
	}
	return res
}

type triplets struct {
	small, middle, large int
}

func newTriplets(s, m, l int) triplets {
	return triplets{s, m, l}
}
