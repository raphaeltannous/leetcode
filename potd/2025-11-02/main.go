func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, 0, m)

	for range m {
		row := make([]int, n)
		grid = append(grid, row)
	}

	result := m * n
	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 3
		result--
	}

	for _, guard := range guards {
		row, column := guard[0], guard[1]
		grid[row][column] = 2
		result--
	}

	for _, guard := range guards {
		row, column := guard[0], guard[1]

		for y := row - 1; y >= 0; y-- {
			if grid[y][column] == 3 || grid[y][column] == 2 {
				break
			}

			if grid[y][column] != 1 {
				result--
			}
			grid[y][column] = 1
		}

		for y := row + 1; y < m; y++ {
			if grid[y][column] == 3 || grid[y][column] == 2 {
				break
			}

			if grid[y][column] != 1 {
				result--
			}
			grid[y][column] = 1
		}

		for x := column - 1; x >= 0; x-- {
			if grid[row][x] == 3 || grid[row][x] == 2 {
				break
			}

			if grid[row][x] != 1 {
				result--
			}
			grid[row][x] = 1
		}

		for x := column + 1; x < n; x++ {
			if grid[row][x] == 3 || grid[row][x] == 2 {
				break
			}

			if grid[row][x] != 1 {
				result--
			}
			grid[row][x] = 1
		}
	}

	return result
}
