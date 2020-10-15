func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	counts := []int{}
	for i := 1; i <= m; i++ {
		if i == 1 {
			blocked := false
			for j := 1; j <= n; j++ {
				if obstacleGrid[i-1][j-1] == 1 {
					blocked = true
				}
				if blocked {
					counts = append(counts, 0)
				} else {
					counts = append(counts, 1)
				}
			}
		} else {
			for j := 1; j <= n; j++ {
				if obstacleGrid[i-1][j-1] == 1 {
					counts[j-1] = 0
					continue
				}

				if j == 1 {
					if counts[0] != 0 {
						counts[0] = 1
					}
				} else {
					counts[j-1] = counts[j-1] + counts[j-2]
				}
			}
		}
	}

	return counts[n-1]
}