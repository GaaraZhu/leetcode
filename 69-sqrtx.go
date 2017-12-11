package main

import "fmt"

func main() {
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return getSqrt(x, x/2, x)
}

func getSqrt(x, left, right int) int {
	if left*left > x {
		return getSqrt(x, left/2, left)
	} else if left*left == x || (left+1)*(left+1) > x {
		return left
	}

	return getSqrt(x, (left+right)/2, right)
}
