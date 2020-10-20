func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	if n == 1 {
		return []string{"()"}
	}

	var previousResult []string
	previousResult = append(previousResult, "()")
	for i := 2; i <= n; i++ {
		tmpResultMap := make(map[string]bool)
		for _, r := range previousResult {
			si := 0
			ei := len(r)
			for ; si <= ei; si++ {
				for ; si <= ei; ei-- {
					newP := r[0:si] + "(" + r[si:ei] + ")" + r[ei:]
					tmpResultMap[newP] = true
				}
			}
		}

		var result []string
		for k, _ := range tmpResultMap {
			result = append(result, k)
		}

		if i == n {
			return result
		} else {
			previousResult = result
		}
	}

	return []string{}
}
