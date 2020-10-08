func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum1 := 0
	sum2 := 0
	for _, num := range nums {
		tmp := sum2
		sum2 = max(sum1+num, sum2)
		sum1 = tmp
	}

	return max(sum1, sum2)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}