func combine(n int, k int) [][]int {
	result := make([][]int, 0)

	subset := make([]int, 0, k)

	var backtrack func(i int)
	backtrack = func(i int) {
		if len(subset) == k {
			newSubset := make([]int, len(subset))
			copy(newSubset, subset)
			result = append(result, newSubset)
			return
		}

		if i > n {
			return
		}

		subset = append(subset, i)
		backtrack(i + 1)

		subset = subset[:len(subset)-1]
		backtrack(i + 1)
	}
	backtrack(1)

	return result
}
