package solution

func orangesRotting(grid [][]int) int {
	q := [][2]int{}
	time, fresh := 0, 0
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	ROWS, COLS := len(grid), len(grid[0])
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if grid[r][c] == 1 {
				fresh += 1
			}
			if grid[r][c] == 2 {
				q = append(q, [2]int{r, c})
			}
		}
	}

	for len(q) > 0 && fresh > 0 {
		length := len(q)
		newQueue := [][2]int{}

		for i := 0; i < length; i++ {
			r, c := q[0][0], q[0][1]
			q = q[1:]
			for _, direction := range directions {
				dr, dc := direction[0], direction[1]
				row, col := dr+r, dc+c

				if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) || grid[row][col] != 1 {
					continue
				}

				grid[row][col] = 2
				newQueue = append(newQueue, [2]int{row, col})
				fresh -= 1
			}
		}
		q = newQueue
		time += 1
	}

	if fresh != 0 {
		return -1
	}

	return time
}
