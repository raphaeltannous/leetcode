func minExtraChar(s string, dictionary []string) int {
	words := make(map[string]bool)
	for _, word := range dictionary {
		words[word] = true
	}

	cache := make(map[int]int)
	cache[len(s)] = 0
	var dfs func(i int) int
	dfs = func(i int) int {
		if val, ok := cache[i]; ok {
			return val
		}

		res := 1 + dfs(i+1)
		for j := i; j < len(s); j++ {
			if words[s[i:j+1]] {
				res = min(res, dfs(j+1))
			}
		}
		cache[i] = res
		return res
	}

	return dfs(0)
}
