func combinationSum(candidates []int, target int) [][]int {
	if target <= 0 {
		return [][]int{}
	}

	result := make([][]int, 0)

	subset := make([]int, 0)
	var backtrack func(i int)
	backtrack = func(i int) {
		sumValue := intSum(subset)
		if sumValue > target {
			return
		}

		if i == len(candidates) {
			if sumValue == target {
				newSubset := make([]int, len(subset))
				copy(newSubset, subset)
				result = append(result, newSubset)
			} else if sumValue < target && sumValue > 0 {
				combinations := combinationSum(subset, target-sumValue)

				for _, comb := range combinations {
					newSubset := make([]int, len(subset))
					copy(newSubset, subset)
					newSubset = append(newSubset, comb...)
					result = append(result, newSubset)
				}
			}

			return
		}

		subset = append(subset, candidates[i])
		backtrack(i + 1)

		subset = subset[:len(subset)-1]
		backtrack(i + 1)
	}
	backtrack(0)

	return result
}

func intSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
