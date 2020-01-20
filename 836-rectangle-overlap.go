package main

import "fmt"

func main() {
	l11 := []int{7, 8, 13, 15}
	l22 := []int{10, 8, 12, 20}
	fmt.Println(isRectangleOverlap(l11, l22))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || rec1[0] >= rec2[2] || rec1[1] >= rec2[3] || rec1[3] <= rec2[1])
}
