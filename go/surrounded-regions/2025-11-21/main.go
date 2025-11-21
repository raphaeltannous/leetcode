func solve(board [][]byte) {
	seenGrid := make([][]bool, len(board))
	for y, row := range board {
		seenGrid[y] = make([]bool, len(row))
	}

	prev := make([][2]int, 0)
	var dfs func(y, x int) bool
	dfs = func(y, x int) bool {
		if y < 0 || y > len(board)-1 || x < 0 || x > len(board[0])-1 {
			return true
		}

		if seenGrid[y][x] {
			return board[y][x] == 'O'
		}

		if board[y][x] == 'X' {
			return false
		}

		prev = append(prev, [2]int{y, x})
		seenGrid[y][x] = true
		board[y][x] = 'X'

		return dfs(y-1, x) || dfs(y+1, x) || dfs(y, x+1) || dfs(y, x-1)
	}

	for y, row := range board {
		for x, _ := range row {
			if seenGrid[y][x] || board[y][x] == 'X' {
				continue
			}

			if dfs(y, x) {
				for _, dir := range prev {
					board[dir[0]][dir[1]] = 'O'
				}
			}

			prev = prev[0:0]
		}
	}
}
