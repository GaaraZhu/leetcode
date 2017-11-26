package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(compareVersion("2.0", "2"))
}

func compareVersion(version1 string, version2 string) int {
	if version1 == "" || version2 == "" {
		return 0
	}

    v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i:=0; i< len(v1); i++ {
		if i > len(v2)-1 {
			if areRestZero(i, v1) {
				return 0
			}
			return 1
		}

		intv, _ := strconv.Atoi(v1[i])
		intv2, _ := strconv.Atoi(v2[i])
		if intv > intv2 {
			return 1
		} else if (intv < intv2) {
			return -1
		}
	}

	if len(v1) < len(v2) {
			if areRestZero(len(v1), v2) {
				return 0
			}
			return -1
	}

	return 0
}

func areRestZero(i int, v []string) bool {
	for ;i<len(v);i++ {
		intv, _ := strconv.Atoi(v[i])
		if intv != 0 {
			return false
		}
	}

	return true
}

