func isAlienSorted(words []string, order string) bool {
	letters := make(map[rune]int)
	for i, letter := range order {
		letters[letter] = i
	}

	p1, p2 := 0, 0
	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i][p1:], words[i+1][p2:]

		if len(word1) == 0 {
			p1, p2 = 0, 0
			continue
		} else if len(word2) == 0 {
			return false
		}

		if result := compareWords(word1, word2, letters); result > 0 {
			return false
		} else if result == 0 {
			p1++
			p2++
			i--
		} else {
			p1, p2 = 0, 0
		}

	}

	return true
}

func compareWords(word1, word2 string, order map[rune]int) int {
	return order[rune(word1[0])] - order[rune(word2[0])]
}
