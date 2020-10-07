func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	result := 0
	for n > 0 {
		n = n / 5
		result += n
	}

	return result
}