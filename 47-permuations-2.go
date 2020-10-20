import "reflect"

func permuteUnique(nums []int) [][]int {
	return permuteFromIndex(0, nums)
}

func permuteFromIndex(index int, nums []int) [][]int {
	var result [][]int
	if index == len(nums)-1 {
		result = append(result, []int{nums[index]})
		return result
	}

	ps := permuteFromIndex(index+1, nums)
	for _, p := range ps {
		for i := 0; i <= len(p); i++ {
			var resultP []int
			for j := 0; j < len(p); j++ {
				if i == j {
					resultP = append(resultP, nums[index])
				}
				resultP = append(resultP, p[j])
			}

			if i == len(p) {
				resultP = append(resultP, nums[index])
			}

			if !checkDuplicated(resultP, result) {
				result = append(result, resultP)
			}
		}
	}

	return result
}

func checkDuplicated(p []int, r [][]int) bool {
	for _, e := range r {
		if reflect.DeepEqual(p, e) {
			return true
		}
	}

	return false
}