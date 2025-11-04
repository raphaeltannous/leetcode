func findXSum(nums []int, k int, x int) []int {
	result := make([]int, 0, len(nums)-k+1)

	subarray := make([]int, 0, k)
	for i := range k {
		subarray = append(subarray, nums[i])
	}

	result = append(result, xSum(subarray, x))
	for i := range len(nums) - k {
		subarray = subarray[1:]
		subarray = append(subarray, nums[k+i])
		result = append(result, xSum(subarray, x))
	}

	return result
}

func xSum(nums []int, x int) int {
	sum := 0

	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	frequencies := make([][2]int, 0, len(freqMap))
	for num, freq := range freqMap {
		frequencies = append(frequencies, [2]int{num, freq})
	}
	slices.SortFunc(frequencies, func(a, b [2]int) int {
		if b[1] == a[1] {
			return b[0] - a[0]
		}

		return b[1] - a[1]
	})

	for i := range x {
		if i < len(frequencies) {
			sum += frequencies[i][0] * frequencies[i][1]
		}
	}

	return sum
}
