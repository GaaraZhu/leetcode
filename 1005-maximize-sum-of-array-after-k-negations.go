import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	if K == 0 {
		return sumArray(A)
	}

	sort.Ints(A)
	for i := 0; i < len(A) && K > 0; i++ {
		if A[i] < 0 {
			A[i] = -1 * A[i]
			K--
		}
	}

	if K == 0 {
		return sumArray(A)
	}

	sort.Ints(A)
	if K%2 != 0 {
		A[0] = -1 * A[0]
	}

	return sumArray(A)
}

func sumArray(A []int) int {
	result := 0
	for _, v := range A {
		result += v
	}

	return result
}