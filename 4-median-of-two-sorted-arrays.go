package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 {
		return findMedian(nums2)
	}
	if n == 0 {
		return findMedian(nums1)
	}
	finalNums := make([]int, 0)
	nums1Index, nums2Index := 0, 0
	for i := 0; i < (m + n); i++ {
		if nums1[nums1Index] < nums2[nums2Index] {
			finalNums = append(finalNums, nums1[nums1Index])
			fmt.Println(finalNums)
			if nums1Index < m-1 {
				nums1Index++
			} else {
				finalNums = append(finalNums, nums2[nums2Index:n]...)
				break
			}
			continue
		}
		finalNums = append(finalNums, nums2[nums2Index])
		if nums2Index < n-1 {
			nums2Index++
		} else {
			finalNums = append(finalNums, nums1[nums1Index:m]...)
			break
		}
	}

	return findMedian(finalNums)
}

func findMedian(nums []int) float64 {
	size := len(nums)
	if size%2 != 0 {
		return float64(nums[size/2])
	}

	return float64((nums[size/2] + nums[size/2-1])) / 2
}
