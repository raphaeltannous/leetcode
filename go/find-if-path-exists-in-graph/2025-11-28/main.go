func validPath(n int, edges [][]int, source int, destination int) bool {
	edgesHash := make(map[int][]int)
	for _, vertices := range edges {
		a, b := vertices[0], vertices[1]
		if _, ok := edgesHash[a]; !ok {
			edgesHash[a] = make([]int, 0)
		}

		if _, ok := edgesHash[b]; !ok {
			edgesHash[b] = make([]int, 0)
		}

		edgesHash[a] = append(edgesHash[a], b)
		edgesHash[b] = append(edgesHash[b], a)
	}

	seen := make(map[int]bool)
	result := false
	var dfs func(edge int)
	dfs = func(edge int) {
		if result {
			return
		}

		if edge == destination {
			result = true
			return
		}

		if seen[edge] {
			return
		}

		es, ok := edgesHash[edge]
		if !ok {
			return
		}

		seen[edge] = true
		for _, e := range es {
			dfs(e)
		}
	}

	dfs(source)
	return result
}
