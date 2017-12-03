package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	var x, y int
	for i:=0; i<len(moves); i++ {
		switch string(moves[i]) {
			case "U": y += 1
			case "D": y -= 1
			case "L": x += 1
			case "R": x -= 1
		}
	}
	return x == 0 && y == 0
}