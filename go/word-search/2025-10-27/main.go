func exist(board [][]byte, word string) bool {
	result := false

	var backtrack func(x, y, i int) bool
	backtrack = func(x, y, i int) bool {
		if i == len(word)-1 {
			return board[y][x] == word[i]
		}

		if board[y][x] == word[i] {
			innerResult := false
			tmp := board[y][x]
			board[y][x] = '.'
			if y > 0 && board[y-1][x] == word[i+1] {
				innerResult = innerResult || backtrack(x, y-1, i+1)
			}

			if y < len(board)-1 && board[y+1][x] == word[i+1] {
				innerResult = innerResult || backtrack(x, y+1, i+1)
			}

			if x > 0 && board[y][x-1] == word[i+1] {
				innerResult = innerResult || backtrack(x-1, y, i+1)
			}

			if x < len(board[y])-1 && board[y][x+1] == word[i+1] {
				innerResult = innerResult || backtrack(x+1, y, i+1)
			}
			board[y][x] = tmp
			return innerResult
		} else {
			return false
		}
	}

RowLoop:
	for y, row := range board {
		for x := range row {
			result = backtrack(x, y, 0)
			if result {
				break RowLoop
			}
		}
	}

	return result
}
