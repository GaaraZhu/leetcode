package main

import "fmt"

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{0, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{0, 1, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 0, 1, 0}))
}

func isOneBitCharacter(bits []int) bool {
	var j = 0
	for j < len(bits)-1 {
		if bits[j] == 0 {
			j++
		} else {
			j += 2
		}
	}

	return j == len(bits)-1
}
