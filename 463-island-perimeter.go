package main

import (
	"fmt"
)

func main() {
	x := [][]int {{0,1,0,0}, {1,1,1,0}, {0,1,0,0},{1,1,0,0}}
	fmt.Println(islandPerimeter(x))
	x = [][]int {{1,0}}
	fmt.Println(islandPerimeter(x))
}

func islandPerimeter(grid [][]int) int {
    var total int
	for i:=0; i<len(grid); i++ {
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] == 1 {
				if i-1<0 || grid[i-1][j] == 0 {
					total = total + 1
				}
				if i+1==len(grid) || grid[i+1][j] == 0 {
					total = total + 1
				}
				if j-1<0 || grid[i][j-1] == 0 {
					total = total + 1
				}
				if j+1==len(grid[i]) ||  grid[i][j+1] == 0 {
					total = total + 1
				}
			}
		}
	}

	return total
}
