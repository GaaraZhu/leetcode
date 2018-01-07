package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(largestNumber([]int{0, 0}))
	fmt.Println(largestNumber([]int{2, 3}))
	fmt.Println(largestNumber([]int{2, 3, 30}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}

func largestNumber(nums []int) string {
	var snums []string
	for _, n := range nums {
		snums = append(snums, strconv.Itoa(n))
	}

	for i:=0; i< len(snums); i++ {
		for j:=i+1; j< len(nums); j++ {
			if snums[i] + snums[j] < snums[j] + snums[i] {
				tmp := snums[i]
				snums[i] = snums[j]
				snums[j] = tmp
			}
		}
	}

	var res string
	var allZero = true
	for _, n := range snums {
		res = res + n
		if n != "0" {
			allZero = false
		}
	}
	if allZero {
		return "0"
	}
	
	return res
}


