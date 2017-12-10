package main

func judgeCircle(moves string) bool {
	var x, y int
	for i := 0; i < len(moves); i++ {
		switch string(moves[i]) {
		case "U":
			y++
		case "D":
			y--
		case "L":
			x++
		case "R":
			x--
		}
	}
	return x == 0 && y == 0
}
