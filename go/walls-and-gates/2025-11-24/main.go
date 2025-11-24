func wallsAndGates(rooms [][]int) {
	var dfs func(y, x, distance int)
	dfs = func(y, x, distance int) {
		if y < 0 || y > len(rooms)-1 || x < 0 || x > len(rooms[0])-1 {
			return
		}

		if rooms[y][x] == 0 || rooms[y][x] == -1 {
			return
		}

		if rooms[y][x] <= distance {
			return
		}

		rooms[y][x] = min(rooms[y][x], distance)

		dfs(y+1, x, distance+1)
		dfs(y-1, x, distance+1)
		dfs(y, x+1, distance+1)
		dfs(y, x-1, distance+1)
	}

	for y, row := range rooms {
		for x, room := range row {
			if room == 0 {
				rooms[y][x] = 1
				dfs(y, x, 0)
			}
		}
	}
}
