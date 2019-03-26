package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	var initString = "1"
	for i := 2; i <= n; i++ {
		initString = count(initString)
	}
	return initString
}

func count(source string) string {
	var buffer bytes.Buffer
	var current string
	var currentCount int
	for i := 0; i < len(source); i++ {
		currentString := string(source[i])
		if currentCount == 0 {
			current = currentString
			currentCount++
		} else {
			if current == currentString {
				currentCount++
			} else {
				buffer.WriteString(strconv.Itoa(currentCount))
				buffer.WriteString(current)
				current = currentString
				currentCount = 1
			}
		}

		if i == len(source)-1 {
			buffer.WriteString(strconv.Itoa(currentCount))
			buffer.WriteString(current)
		}
	}
	return buffer.String()
}
