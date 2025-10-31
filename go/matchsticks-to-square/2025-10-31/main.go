func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	matchsticksSum := sIntSum(matchsticks)
	if matchsticksSum%4 != 0 {
		return false
	}

	slices.SortFunc(matchsticks, func(a, b int) int {
		return b - a
	})

	wantedSum := matchsticksSum / 4

	used := make([]bool, len(matchsticks))
	var backtrack func(i, subsetSum, sidesLeft int) bool
	backtrack = func(i, subsetSum, sidesLeft int) bool {
		if sidesLeft == 0 {
			return true
		}

		if subsetSum == wantedSum {
			return backtrack(0, 0, sidesLeft-1)
		}

		for j := i; j < len(matchsticks); j++ {
			if used[j] || subsetSum+matchsticks[j] > wantedSum {
				continue
			}

			used[j] = true
			if backtrack(j+1, subsetSum+matchsticks[j], sidesLeft) {
				return true
			}
			used[j] = false
		}
		return false
	}

	return backtrack(0, 0, 4)
}

func sIntSum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
