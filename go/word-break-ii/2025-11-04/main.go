func wordBreak(s string, wordDict []string) []string {
	result := make([]string, 0)

	currentString := make([]string, 0, len(wordDict))
	var backtrack func()
	backtrack = func() {
		if len(s) == 0 {
			result = append(result, strings.Join(currentString, " "))
			return
		}

		for _, word := range wordDict {
			if len(word) <= len(s) {
				substring := s[:len(word)]
				if substring == word {
					s = s[len(word):]
					currentString = append(currentString, word)
					backtrack()
					s = word + s
					currentString = currentString[:len(currentString)-1]
				}
			}
		}
	}
	backtrack()

	return result
}
