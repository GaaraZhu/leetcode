package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStrings2("0", "0"))
	fmt.Println(addStrings2("10", "0"))
	fmt.Println(addStrings2("0", "1110"))
	fmt.Println(addStrings2("999991", "19"))
	fmt.Println(addStrings2("2221", "1"))
	fmt.Println(addStrings2("2221", "2221"))
}

func addStrings2(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return addStringsForOrderedNum(num1, num2)
	}
	return addStringsForOrderedNum(num2, num1)
}

func addStringsForOrderedNum(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	var res string
	var addOneToLeft bool
	for i := 0; i <= l1-1; i++ {
		n1, err := strconv.Atoi(num1[l1-1-i : l1-i])
		if err != nil {
			panic("something is wrong!!")
		}
		n2, err := strconv.Atoi(num2[l2-1-i : l2-i])
		if err != nil {
			panic("something is wrong!!!")
		}
		var value = n1 + n2
		if addOneToLeft {
			value++
		}

		if value > 9 {
			addOneToLeft = true
			value = value % 10
		} else {
			addOneToLeft = false
		}
		res = fmt.Sprintf("%v", value) + res
	}

	for i := l1; i <= l2-1; i++ {
		if addOneToLeft {
			n1, err := strconv.Atoi(num2[l2-1-i : l2-i])
			if err != nil {
				panic("something is wrong!!!!")
			}
			var value = n1 + 1
			if value > 9 {
				addOneToLeft = true
				value = 0
			} else {
				addOneToLeft = false
			}
			res = fmt.Sprintf("%v", value) + res
		} else {
			res = num2[l2-1-i:l2-i] + res
		}
	}

	if addOneToLeft {
		res = "1" + res
	}

	return res
}

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)

	var size = l2
	if l1 < l2 {
		num1 = appendZeroAhead(num1, l2-l1)
	} else if l2 < l1 {
		size = l1
		num2 = appendZeroAhead(num2, l1-l2)
	}

	var addOneToLeft bool
	var result = make([]int, size+1)

	for i := 0; i <= size-1; i++ {
		intNum1, err := strconv.Atoi(num1[size-1-i : size-i])
		if err != nil {
			panic("Something is wrong!!")
		}
		intNum2, err := strconv.Atoi(num2[size-1-i : size-i])
		if err != nil {
			panic("Something is wrong!!!")
		}

		result[size-i] = intNum1 + intNum2

		if addOneToLeft {
			result[size-i] = result[size-i] + 1
		}
		if result[size-i] > 9 {
			addOneToLeft = true
			result[size-i] = result[size-i] % 10
		} else {
			addOneToLeft = false
		}
	}

	if addOneToLeft {
		result[0] = 1
	}

	return intArrayToString(result)
}

func appendZeroAhead(num string, numOfZeros int) string {
	var zeros string
	for i := 0; i < numOfZeros; i++ {
		zeros += "0"
	}
	return zeros + num
}

func intArrayToString(nums []int) string {
	var result string
	for i, num := range nums {
		if i == 0 && num == 0 {
			continue
		}
		result = result + fmt.Sprintf("%v", num)
	}

	return result
}
