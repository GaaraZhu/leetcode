package main

import (
	"fmt"
)

func main() {
	fmt.Println(addDigits(5))
	fmt.Println(addDigits(15))
	fmt.Println(addDigits(165))
	fmt.Println(addDigits(55))
	fmt.Println(addDigits(9999))

	fmt.Println(addDigits2(5))
	fmt.Println(addDigits2(15))
	fmt.Println(addDigits2(165))
	fmt.Println(addDigits2(55))
	fmt.Println(addDigits2(9999))
}

func addDigits(num int) int {
	if num/10 == 0 {
		return num
	}

	return addDigits(addDigits(num/10) + addDigits(num%10))
}

func addDigits2(num int) int {
	if num == 0 {
		return 0
	}

	if num%9 == 0 {
		return 9
	}
	return num % 9
}
