func maxAreaOfIsland(grid [][]int) int {
	var maxArea int
	var area int

	var dfs func(y, x int)
	dfs = func(y, x int) {
		if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1 || grid[y][x] == 0 {
			return
		}

		area++
		grid[y][x] = 0
		dfs(y+1, x)
		dfs(y-1, x)
		dfs(y, x+1)
		dfs(y, x-1)
	}

	for y, row := range grid {
		for x, cell := range row {
			if cell == 1 {
				area = 0
				dfs(y, x)
				maxArea = max(maxArea, area)
			}
		}
	}

	return maxArea
}
