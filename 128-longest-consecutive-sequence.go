package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{}))
	fmt.Println(longestConsecutive([]int{0}))
	fmt.Println(longestConsecutive([]int{0, 0}))
	fmt.Println(longestConsecutive([]int{0, -1}))
	fmt.Println(longestConsecutive([]int{0, -1, -2, 4, 5, 6, 1}))
	fmt.Println(longestConsecutive([]int{100, 99, 6, 4, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{2147483646, 2147483644, 2147483645, 0}))
	fmt.Println(longestConsecutive([]int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645}))
}

func longestConsecutive(nums []int) (max int) {
	m := map[int]interface{}{}
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for k := range m {
		var currentMax int
		currentMax++
		delete(m, k)
		for j := k - 1; ; j-- {
			_, ok := m[j]
			if !ok {
				break
			}
			currentMax++
			delete(m, j)
		}
		for j := k + 1; ; j++ {
			_, ok := m[j]
			if !ok {
				break
			}
			currentMax++
			delete(m, j)
		}
		if currentMax > max {
			max = currentMax
		}
	}
	return
}
