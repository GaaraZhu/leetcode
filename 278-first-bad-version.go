func firstBadVersion(n int) int {
	return findFirstBadVersion(n, n)
}

func findFirstBadVersion(i int, last int) int {
	if isBadVersion(i) {
		if !isBadVersion(i - 1) {
			return i
		} else {
			return findFirstBadVersion(i/2, i)
		}
	} else {
		if isBadVersion(i + 1) {
			return i + 1
		}
		return findFirstBadVersion(i+(last-i)/2, last)
	}
}