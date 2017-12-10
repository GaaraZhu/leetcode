package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("", "11"))
	fmt.Println(addBinary("10", ""))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("11", "111"))
}

func addBinary(a string, b string) string {
	if a == "" {
		return b
	}

	if b == "" {
		return a
	}

	var aa, ba []int
	if len(a) > len(b) {
		for i := 0; i < len(a)-len(b); i++ {
			ba = append(ba, 0)
		}
		ba = append(ba, stringToArray(b)...)
		aa = stringToArray(a)
	} else {
		for i := 0; i < len(b)-len(a); i++ {
			aa = append(aa, 0)
		}
		aa = append(aa, stringToArray(a)...)
		ba = stringToArray(b)
	}

	var res = make([]int, len(aa)+1)
	var addOneToLeft bool
	for i := len(aa) - 1; i >= 0; i-- {
		if aa[i] == 1 && ba[i] == 1 {
			if addOneToLeft {
				res[i+1] = 1
			} else {
				res[i+1] = 0
			}
			addOneToLeft = true
		} else if aa[i] == 1 || ba[i] == 1 {
			if addOneToLeft {
				res[i+1] = 0
				addOneToLeft = true
			} else {
				res[i+1] = 1
				addOneToLeft = false
			}
		} else {
			if addOneToLeft {
				res[i+1] = 1
			} else {
				res[i+1] = 0
			}
			addOneToLeft = false
		}
	}
	if addOneToLeft {
		res[0] = 1
	}

	return arrayToString(res)
}

func stringToArray(s string) []int {
	var res []int
	for _, r := range s {
		v, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		res = append(res, v)
	}

	return res
}

func arrayToString(a []int) string {
	if len(a) == 0 {
		return ""
	}
	var res string
	for i, v := range a {
		if i == 0 && v == 0 {
			continue
		}
		res = res + strconv.Itoa(v)
	}

	return res
}
