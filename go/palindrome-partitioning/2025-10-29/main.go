func partition(s string) [][]string {
	result := make([][]string, 0)

	subset := make([]string, 0, len(s))
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(s) {
			isValid := true
			for _, word := range subset {
				isValid = isValid && isPalindrome(word)
			}

			if isValid {
				newSubset := make([]string, len(subset))
				copy(newSubset, subset)
				result = append(result, newSubset)
			}
			return
		}

		if len(subset) > 0 {
			subset = append(subset, string(s[i]))
			backtrack(i + 1)

			subset = subset[:len(subset)-1]
			tmp := subset[len(subset)-1] + string(s[i])
			subset = subset[:len(subset)-1]
			subset = append(subset, tmp)
			backtrack(i + 1)
		} else {
			subset = append(subset, string(s[i]))
			backtrack(i + 1)
		}
	}
	backtrack(0)

	return result
}

func isPalindrome(word string) bool {
	j := len(word) - 1
	for x := range word {
		if word[x] != word[j] {
			return false
		}

		if x == j {
			break
		}

		j--
	}

	return true
}
