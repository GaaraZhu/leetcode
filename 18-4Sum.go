package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{}, 0))                      //[]
	fmt.Println(fourSum([]int{1, 2}, 3))                  //[]
	fmt.Println(fourSum([]int{-1, 0, 1, 2, -1, -4}, 0))   //[[-1 0 1] [-1 -1 2]]
	fmt.Println(fourSum([]int{0, 0, 0, 1, 1, -1, -2}, 1)) //[[0 0 0] [-1 0 1] [-2 1 1]]
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	resMap := map[quadruplets]struct{}{}
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		if i > len(nums)-3 {
			break
		} else if i == len(nums)-4 && nums[i+1]+nums[i+2]+nums[i+3] == t {
			resMap[newQuadruplets(nums[i], nums[i+1], nums[i+2], nums[i+3])] = struct{}{}
		}
		quadruplets := restThreeSum(nums[i+1:len(nums)], nums[i], t)
		for k := range quadruplets {
			resMap[k] = struct{}{}
		}
	}

	var res = [][]int{}
	for k := range resMap {
		res = append(res, []int{k.ss, k.small, k.middle, k.large})
	}
	return res
}

func restThreeSum(nums []int, start, target int) map[quadruplets]struct{} {
	quadrupletsMap := map[quadruplets]struct{}{}
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		var j = i + 1
		var k = len(nums) - 1

		for j < k {
			switch {
			case nums[j]+nums[k] == t:
				quadrupletsMap[newQuadruplets(start, nums[i], nums[j], nums[k])] = struct{}{}
				j++
			case nums[j]+nums[k] > t:
				k--
			default:
				j++
			}
		}
	}

	return quadrupletsMap
}

type quadruplets struct {
	ss, small, middle, large int
}

func newQuadruplets(ss, s, m, l int) quadruplets {
	return quadruplets{ss, s, m, l}
}
