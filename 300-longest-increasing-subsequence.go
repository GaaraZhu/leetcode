func lengthOfLIS(nums []int) int {
	maxUint := ^uint(0)
	maxInt := int(maxUint >> 1)
	minInt := -1*maxInt - 1

	memo := make(map[key]int)
	return lengthOfPosition(nums, minInt, 0, memo)
}

func lengthOfPosition(nums []int, previous int, currentIndex int, memo map[key]int) int {
	key := key{previous: previous, currentIndex: currentIndex}
	if v, ok := memo[key]; ok {
		return v
	}

	if currentIndex == len(nums) {
		return 0
	}

	lengthWhenIncludeCurrent := 0
	if nums[currentIndex] > previous {
		lengthWhenIncludeCurrent = lengthOfPosition(nums, nums[currentIndex], currentIndex+1, memo) + 1
	}

	lengthWhenNotIncludeCurrent := lengthOfPosition(nums, previous, currentIndex+1, memo)

	result := max(lengthWhenIncludeCurrent, lengthWhenNotIncludeCurrent)
	memo[key] = result
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

type key struct {
	previous     int
	currentIndex int
}