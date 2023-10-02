package number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	islands := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				bfs(grid, row, col)
				islands++
			}
		}
	}

	return islands
}

func bfs(grid [][]byte, row, col int) {
	rows, cols := len(grid), len(grid[0])
	queue := [][]int{{row, col}}

	for len(queue) > 0 {
		currRow, currCol := queue[0][0], queue[0][1]
		queue = queue[1:]

		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, direction := range directions {
			newRow, newCol := currRow+direction[0], currCol+direction[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == '1' {
				grid[newRow][newCol] = '2'
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
}
