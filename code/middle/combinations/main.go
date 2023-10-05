package combinations

func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	if n < k {
		return result
	}

	backtrack(1, n, k, []int{}, &result)

	return result
}

func backtrack(start, n, k int, combinations []int, result *[][]int) {
	if len(combinations) == k {
		*result = append(*result, append([]int{}, combinations...))
		return
	}

	for i := start; i <= n; i++ {
		combinations = append(combinations, i)
		backtrack(i+1, n, k, combinations, result)
		combinations = combinations[:len(combinations)-1]
	}
}
