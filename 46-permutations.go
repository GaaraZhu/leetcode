func permute(nums []int) [][]int {
	return perMuteFromIndex(0, nums)
}

func perMuteFromIndex(index int, nums []int) [][]int {
	var result [][]int
	if index == len(nums)-1 {
		result = append(result, []int{nums[index]})
		return result
	}

	ps := perMuteFromIndex(index+1, nums)
	for _, p := range ps {
		for i := 0; i <= len(p); i++ {
			resultP := []int{}
			for j := 0; j < len(p); j++ {
				if i == j {
					resultP = append(resultP, nums[index])
				}
				resultP = append(resultP, p[j])
			}

			if i == len(p) {
				resultP = append(resultP, nums[index])
				result = append(result, resultP)
				break
			}
			result = append(result, resultP)
		}
	}

	return result
}
