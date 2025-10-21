func subsetXORSum(nums []int) int {
	result := 0

	var backtrack func(prevXOR int, numbers []int)
	backtrack = func(prevXOR int, numbers []int) {
		if len(numbers) == 0 {
			result += prevXOR
			return
		}

		backtrack(prevXOR, numbers[1:])
		backtrack(prevXOR^numbers[0], numbers[1:])
	}
	backtrack(0, nums)

	return result
}
