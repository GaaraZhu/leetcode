func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := []int{}
	dp = append(dp, 1)
	maxAns := 1
	for i := 1; i < len(nums); i++ {
		maxVal := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxVal = max(maxVal, dp[j])
			}
		}
		dp = append(dp, maxVal+1)
		maxAns = max(maxAns, dp[i])
	}

	return maxAns
}

func max(a int, b int) int {
	if a < b {
		return b
	}

	return a
}
