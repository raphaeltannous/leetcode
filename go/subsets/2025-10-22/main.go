import "math"

func subsets(nums []int) [][]int {
	result := make([][]int, 0, int(math.Pow(2, float64(len(nums)))))

	var backtrack func(prev, numbers []int)
	backtrack = func(prev, numbers []int) {
		if len(numbers) == 0 {
			result = append(result, prev)
			return
		}

		backtrack(prev, numbers[1:])

		newPrev := make([]int, len(prev))
		copy(newPrev, prev)
		newPrev = append(newPrev, numbers[0])
		backtrack(newPrev, numbers[1:])
	}
	backtrack([]int{}, nums)

	return result
}
