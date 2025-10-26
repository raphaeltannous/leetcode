func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	slices.Sort(nums)

	result := make([][]int, 0)

	subset := make([]int, 0, len(nums))
	var backtrack func(i int)
	backtrack = func(i int) {
		if len(subset) == len(nums) {
			combs := permuteUnique(subset[1:])

			if len(combs) > 0 {
				for _, comb := range combs {
					newSubset := make([]int, 1)
					copy(newSubset, subset)
					newSubset = append(newSubset, comb...)
					result = append(result, newSubset)
				}
			} else {
				newSubset := make([]int, len(subset))
				copy(newSubset, subset)
				result = append(result, newSubset)
			}

			return
		}

		if i >= len(nums) {
			return
		}

		subset = append(subset, nums[i])
		backtrack(i + 1)

		subset = subset[:len(subset)-1]
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
		backtrack(i + 1)
	}

	for x, num := range nums {
		if x > 0 && num == nums[x-1] {
			continue
		}
		nums[0], nums[x] = nums[x], nums[0]
		backtrack(0)
		nums[0], nums[x] = nums[x], nums[0]
	}

	return result
}
