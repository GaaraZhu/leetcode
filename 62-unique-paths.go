// vision 1: recursion
func uniquePaths(m int, n int) int {
	if (m == 1 && n == 2) || (m == 2 && n == 1) || (m == 1 && n == 1) {
		return 1
	} else if m >= 2 && n >= 2 {
		return uniquePaths(m, n-1) + uniquePaths(m-1, n)
	} else if m >= 2 {
		return uniquePaths(m-1, n)
	} else if n >= 2 {
		return uniquePaths(m, n-1)
	}

	return 0
}

// vision 2: recursion with memo
func uniquePaths(m int, n int) int {
	memo := make(map[location]int)
	return uniquePathsWithMemo(m, n, memo)
}

func uniquePathsWithMemo(m int, n int, memo map[location]int) int {
	if v, ok := memo[location{m: m, n: n}]; ok {
		return v
	}
	result := 0
	defer func() {
		memo[location{m: m, n: n}] = result
	}()

	if (m == 1 && n == 2) || (m == 2 && n == 1) || (m == 1 && n == 1) {
		result = 1
	} else if m >= 2 && n >= 2 {
		result = uniquePaths(m, n-1) + uniquePaths(m-1, n)
	} else if m >= 2 {
		result = uniquePaths(m-1, n)
	} else if n >= 2 {
		result = uniquePaths(m, n-1)
	}

	return result
}

type location struct {
	m int
	n int
}

// vision 3: dynamic programming
func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	counts := []int{}
	for i := 1; i <= m; i++ {
		if i == 1 {
			for j := 1; j <= n; j++ {
				counts = append(counts, 1)
			}
			continue
		} else {
			for j := 1; j <= n; j++ {
				if j == 1 {
					counts[0] = 1
					continue
				}
				counts[j-1] = counts[j-1] + counts[j-2]
			}
		}
	}

	return counts[n-1]
}
