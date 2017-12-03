package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(9))
}

func isPowerOfThree(n int) bool {
	return n>0 && 1162261467%n==0
}
