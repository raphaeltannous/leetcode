func orangesRotting(grid [][]int) int {
	result := 0

	seenGrid := make([][]bool, len(grid))
	for x, row := range grid {
		seenGrid[x] = make([]bool, len(row))
	}

	queue := make([][3]int, 0)

	for y, row := range grid {
		for x, orange := range row {
			if orange == 2 {
				queue = append(queue, [3]int{y, x, 0})
			}
		}
	}

	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]

		y, x, minute := front[0], front[1], front[2]

		if seenGrid[y][x] || grid[y][x] == 0 {
			continue
		}

		seenGrid[y][x] = true
		result = max(minute, result)

		dirs := [][2]int{
			{y + 1, x},
			{y - 1, x},
			{y, x + 1},
			{y, x - 1},
		}

		for _, dir := range dirs {
			if rotOrange(grid, dir[0], dir[1]) {
				queue = append(queue, [3]int{dir[0], dir[1], minute + 1})
			}
		}
	}

	for _, row := range grid {
		for _, orange := range row {
			if orange == 1 {
				return -1
			}
		}
	}

	return result
}

func rotOrange(grid [][]int, y, x int) bool {
	if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1 {
		return false
	}

	if grid[y][x] == 1 {
		grid[y][x] = 2
		return true
	}

	return false
}
