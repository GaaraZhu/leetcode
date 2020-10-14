// first version: recursion
func canJump(nums []int) bool {
	return jump(nums, 0)
}

func jump(nums []int, i int) bool {
	if len(nums)-1-i <= nums[i] {
		return true
	}

	for j := 1; j <= nums[i]; i++ {
		if jump(nums, i+j) {
			return true
		}
	}
	return false
}

// second version: recursion with memo(top-down)
func canJump(nums []int) bool {
	memo := make(map[int]bool)
	return jump(nums, 0, memo)
}

func jump(nums []int, i int, memo map[int]bool) bool {
	if v, ok := memo[i]; ok {
		return v
	}

	if len(nums)-1-i <= nums[i] {
		memo[i] = true
		return true
	}

	for j := 1; j <= nums[i] && j <= len(nums)-1; j++ {
		if jump(nums, i+j, memo) {
			memo[i] = true
			return true
		}
	}
	memo[i] = false
	return false
}

// third version: recusion-less(bottom-up)
func canJump(nums []int) bool {
	memo := make(map[int]bool)
	memo[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		// check numbers after index i to see whether there one which can jump to the end
		for j := i + 1; j <= len(nums)-1 && j <= nums[i]+i; j++ {
			if v, ok := memo[j]; ok && v {
				memo[i] = true
				break
			}
		}
	}

	return memo[0]
}

// fourth version: Greedy
func canJump(nums []int) bool {

	lastGoodNumber := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		// check whether current number can jump to the last good number which can jump to the end
		if i+nums[i] >= lastGoodNumber {
			lastGoodNumber = i
		}
	}

	return lastGoodNumber == 0
}
