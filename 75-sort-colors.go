func sortColors(nums []int) {
	colorMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if c, ok := colorMap[nums[i]]; ok {
			colorMap[nums[i]] = c + 1
		} else {
			colorMap[nums[i]] = 1
		}
	}

	index := 0
	if c, ok := colorMap[0]; ok {
		for i := 0; i < c; i++ {
			nums[index] = 0
			index++
		}
	}

	if c, ok := colorMap[1]; ok {
		for i := 0; i < c; i++ {
			nums[index] = 1
			index++
		}
	}

	if c, ok := colorMap[2]; ok {
		for i := 0; i < c; i++ {
			nums[index] = 2
			index++
		}
	}
}