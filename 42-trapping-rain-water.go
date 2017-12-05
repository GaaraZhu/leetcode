package main

import (
    "fmt"
)

func main() {
    input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    fmt.Println(trap(input))

    input = []int{8, 5, 4, 1, 8, 9, 3, 0, 0}
    fmt.Println(trap(input))
}

func trap(height []int) int {
    return trapSingle(height)
}

func trapSingle(height []int) int {
    if len(height) < 3 {
        return 0
    }

    var top, secondTop int
    var topIndex, secondTopIndex int
    for i := 0; i < len(height); i++ {
        if height[i] > top {
            secondTop = top
            secondTopIndex = topIndex

            top = height[i]
            topIndex = i
        } else if topIndex == secondTopIndex || height[i] > secondTop {
            secondTop = height[i]
            secondTopIndex = i
        }
    }

    var smallerIndex = secondTopIndex
    var biggerIndex = topIndex
    if secondTopIndex > topIndex {
        smallerIndex = topIndex
        biggerIndex = secondTopIndex
    }
    var size = secondTop * (biggerIndex - smallerIndex - 1)
    for i := smallerIndex + 1; i < biggerIndex; i++ {
        size = size - height[i]
    }

    return size + trapSingle(height[0:smallerIndex+1]) + trapSingle(height[biggerIndex:len(height)])
}