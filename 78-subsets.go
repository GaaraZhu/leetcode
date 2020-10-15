func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})

	for i := 0; i < len(nums); i++ {
		var newResult [][]int
		for _, r := range result {
			subset := []int{} // cannot use something like: subset := (r, nums[i]) because that will change r
			for _, v := range r {
				subset = append(subset, v)
			}
			subset = append(subset, nums[i])
			newResult = append(newResult, subset)
		}

		for _, r := range newResult {
			result = append(result, r)
		}
	}

	return result
}