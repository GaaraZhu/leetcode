package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{}))                      //[]
	fmt.Println(quickSort([]int{1, 2}))                  //[1, 2]
	fmt.Println(quickSort([]int{1, 2, 0}))               //[0, 1, 2]
	fmt.Println(quickSort([]int{-1, 0, 1, 2, -1, -4}))   //[-4, -1, -1, 0, 1, 2]
	fmt.Println(quickSort([]int{0, 0, 0, 1, 1, -1, -2})) //[-2, -1, 0, 0, 0, 1, 1]
}

func quickSort(nums []int) []int {
	return sort(nums, 0, len(nums)-1)
}

func sort(nums []int, left, right int) []int {
	if left >= right {
		return nums
	}

	tmp := nums[left]
	var i = left
	var j = right
	for i < j {
		for nums[j] >= tmp && i < j {
			j--
		}
		for nums[i] <= tmp && i < j {
			i++
		}
		swap(nums, i, j)
	}

	swap(nums, left, i)

	sort(nums, left, i-1)
	sort(nums, i+1, right)

	return nums
}

func swap(nums []int, i, j int) {
	if i >= len(nums) || j >= len(nums) {
		return
	}
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}
