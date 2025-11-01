func canPartitionKSubsets(nums []int, k int) bool {
	numsSum := sIntSum(nums)

	if numsSum%k != 0 {
		return false
	}

	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})

	wantedSum := numsSum / k
	isUsed := make([]bool, len(nums))
	var backtrack func(i, kLeft, currentSum int) bool
	backtrack = func(i, kLeft, currentSum int) bool {
		if kLeft == 0 {
			return true
		}

		if currentSum == wantedSum {
			return backtrack(0, kLeft-1, 0)
		}

		for j := i; j < len(nums); j++ {
			if isUsed[j] || currentSum+nums[j] > wantedSum {
				continue
			}

			isUsed[j] = true
			if backtrack(j+1, kLeft, currentSum+nums[j]) {
				return true
			}
			isUsed[j] = false
		}

		return false
	}

	return backtrack(0, k, 0)
}

func sIntSum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
