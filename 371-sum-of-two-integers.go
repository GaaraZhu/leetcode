package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSum(25, 8))
}

func getSum(a int, b int) int {
    nc := a ^ b
    c := a & b << 1
    if (c ==0) {
	return nc | c
    }
    
   return getSum(nc, c)
}

