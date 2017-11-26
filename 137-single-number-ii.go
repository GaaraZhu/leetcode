package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{-2,-2,1,1,-3,1,-3,-3,-4,-2}))
}

func singleNumber(nums []int) int {
  	var (
        bits     uint   = 64
        i64      int64 = 1 << 32
    )
	if int(i64) == 0 {
        bits = 32
    }
      var res uint
	  var i uint = 0
	  for ; i<bits; i++ {
		  var sum int = 0
		  for j:=0; j<len(nums); j++ {
			  sum += (nums[j]>>i) & 1
		  }
		  res = res | uint(sum % 3) << i
	  }

	  return int(res)
}
