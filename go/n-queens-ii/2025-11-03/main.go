func totalNQueens(n int) int {
	result := 0

	current := make([][]int, 0, n)
	for range n {
		current = append(current, make([]int, n))
	}

	var backtrack func(x, nLeft int)
	backtrack = func(x, nLeft int) {
		if nLeft == 0 {
			result += 1
			return
		}

		if x > n {
			return
		}

		skipped := 0
		for i, cell := range current[x] {
			if cell == 0 {
				current[x][i] = -1
				setCellsStatus(x, i, current, true)
				backtrack(x+1, nLeft-1)
				current[x][i] = 0
				setCellsStatus(x, i, current, false)
			} else {
				skipped++

				if skipped == n {
					return
				}
			}
		}
	}
	backtrack(0, n)

	return result
}

func setCellsStatus(y, x int, grid [][]int, status bool) {
	cells := [][2]int{
		{y - 1, x - 1},
		{y - 1, x},
		{y - 1, x + 1},
		{y, x - 1},
		{y, x + 1},
		{y + 1, x - 1},
		{y + 1, x},
		{y + 1, x + 1},
	}

	for _, cell := range cells {
		yOperation := cell[0] - y
		xOperation := cell[1] - x

		for cy, cx := y+yOperation, x+xOperation; cy >= 0 && cy < len(grid) && cx >= 0 && cx < len(grid); cy, cx = cy+yOperation, cx+xOperation {
			setCellStatus(cy, cx, grid, status)
		}
	}
}

func setCellStatus(y, x int, grid [][]int, status bool) {
	if y < 0 || x < 0 || y > len(grid) || x > len(grid) {
		return
	}

	if status {
		grid[y][x]++
	} else {
		grid[y][x]--
	}
}
