func pacificAtlantic(heights [][]int) [][]int {
	rowsLen, colsLen := len(heights), len(heights[0])

	pacificConnection := make([][]bool, rowsLen)
	atlanticConnection := make([][]bool, rowsLen)
	for y := range heights {
		pacificConnection[y] = make([]bool, colsLen)
		atlanticConnection[y] = make([]bool, colsLen)
	}

	pacificQueue := make([][3]int, 0)
	atlanticQueue := make([][3]int, 0)
	for y := range rowsLen {
		pacificQueue = append(pacificQueue, [3]int{y, 0, -1})
		atlanticQueue = append(atlanticQueue, [3]int{y, colsLen - 1, -1})
	}
	for x := range colsLen {
		pacificQueue = append(pacificQueue, [3]int{0, x, -1})
		atlanticQueue = append(atlanticQueue, [3]int{rowsLen - 1, x, -1})
	}

	bfs(heights, pacificConnection, pacificQueue)
	bfs(heights, atlanticConnection, atlanticQueue)

	result := make([][]int, 0)

	for y, row := range heights {
		for x := range row {
			if pacificConnection[y][x] && atlanticConnection[y][x] {
				result = append(result, []int{y, x})
			}
		}
	}

	return result
}

func bfs(grid [][]int, connectionGrid [][]bool, queue [][3]int) {
	for len(queue) != 0 {
		height := queue[0]
		queue = queue[1:]

		y, x, prev := height[0], height[1], height[2]

		if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1 {
			continue
		}

		if connectionGrid[y][x] {
			continue
		}

		curr := grid[y][x]

		if curr < prev {
			continue
		}

		connectionGrid[y][x] = true

		queue = append(queue, [3]int{y + 1, x, curr})
		queue = append(queue, [3]int{y - 1, x, curr})
		queue = append(queue, [3]int{y, x + 1, curr})
		queue = append(queue, [3]int{y, x - 1, curr})
	}
}
