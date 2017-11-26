package main

import (
	"fmt"
)

func main() {
	a := []int{2,3,5,3,2,6,7,6,5}
	fmt.Println(singleNumber(a))
}

func singleNumber(nums []int) int {
    var b = 0
    for _, n := range nums {
        b = b ^ n
    }
    return b
}
