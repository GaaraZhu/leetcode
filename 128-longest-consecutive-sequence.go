package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{}))
	fmt.Println(longestConsecutive([]int{0}))
	fmt.Println(longestConsecutive([]int{0, 0}))
	fmt.Println(longestConsecutive([]int{0, -1}))
	fmt.Println(longestConsecutive([]int{0, -1, -2, 4, 5, 6, 1}))
	fmt.Println(longestConsecutive([]int{100, 200, 7, 6, 4, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	var positives, negatives []int
	var containsZero bool
	for _, num := range nums {
		switch {
		case num == 0:
			containsZero = true
		case num > 0:
			positives = append(positives, num)
		case num < 0:
			negatives = append(negatives, 0 - num)
		}
	}

	var longestInPos, longestInNeg, smallestValueInPos, smallestValueInNeg int
	longestInPos, smallestValueInPos = findLongestConsecutive(positives)
	longestInNeg, smallestValueInNeg = findLongestConsecutive(negatives)
	if containsZero {
		if longestInPos == 0 && longestInNeg == 0 {
			return 1
		}
		if smallestValueInPos == 1 && smallestValueInNeg == 1 {
			return longestInPos + 1 + longestInNeg
		}
		if smallestValueInPos == 1 {
			longestInPos += 1
		}
		if smallestValueInNeg == 1 {
			longestInNeg += 1
		}
	}
	if longestInPos > longestInNeg {
		return longestInPos
	} else {
		return longestInNeg
	}
}


func findLongestConsecutive(nums []int) (int, int) {
    if len(nums) == 0 {
		return 0, 0
	}

	var rs = uint(0)
	var biggest int
	for _, num := range nums {
		if num > biggest {
			biggest = num
		}
		rs = rs | (1 << uint(num))
	}

	var lastMax, lastValue, smallestOne int
	var currentTotal = 1
	for i:=1; i<=biggest; i++ {
		currentValue := int((rs >> uint(i)) & 1)
		if currentValue == 1 {
			if lastValue != 1 {
				smallestOne = i
			} else {
				currentTotal += 1
			}
		} 

		if currentValue != 1 && lastMax < currentTotal {
			lastMax = currentTotal
			currentTotal = 1
		}

		lastValue = currentValue
	}

	if lastMax < currentTotal {
		return currentTotal, smallestOne
	}

	return lastMax, smallestOne
}