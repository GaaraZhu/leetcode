func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}

	current := newInterval
	for i := 0; i < len(intervals); i++ {
		if current == nil || len(current) != 2 {
			result = append(result, intervals[i])
			continue
		}

		if intervals[i][0] > current[1] {
			result = append(result, current)
			result = append(result, intervals[i])
			current = nil
			continue
		} else if intervals[i][1] < current[0] {
			result = append(result, intervals[i])
			if i == len(intervals)-1 {
				result = append(result, current)
			}
			continue
		} else {
			left := min(intervals[i][0], current[0])
			right := max(intervals[i][1], current[1])

			if i+1 < len(intervals) {
				if intervals[i+1][0] > right {
					result = append(result, []int{left, right})
					current = nil
				} else {
					current = []int{left, right}
				}
			} else {
				result = append(result, []int{left, right})
			}
		}
	}

	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}