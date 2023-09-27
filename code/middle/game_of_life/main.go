package game_of_life

func gameOfLife(board [][]int) {
	// 0 -> 0 = 0
	// 0 -> 1 = 2
	// 1 -> 0 = 3
	// 1 -> 1 = 1
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			liveNeighbours := countLiveNeighbours(board, i, j)
			if board[i][j] == 1 && (liveNeighbours < 2 || liveNeighbours > 3) {
				board[i][j] = 3
			}
			if board[i][j] == 0 && liveNeighbours == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] >>= 1
		}
	}
}

func countLiveNeighbours(board [][]int, i, j int) int {
	liveNeighbours := 0
	for x := max(0, i-1); x <= min(len(board)-1, i+1); x++ {
		for y := max(0, j-1); y <= min(len(board[0])-1, j+1); y++ {
			if x == i && y == j {
				continue
			}
			if board[x][y] == 1 || board[x][y] == 3 {
				liveNeighbours++
			}
		}
	}
	return liveNeighbours
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gameOfLifeNeetCode(board [][]int) {
	/*
		Coding the rules in-place:
		- 0 -> 0 = 0
		- 1 -> 0 = 1
		- 0 -> 1 = 2
		- 1 -> 1 = 3
	*/

	rows, cols := len(board), len(board[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNeighbours := countLiveNeighbours(board, row, col)

			if board[row][col] == 1 && (liveNeighbours == 2 || liveNeighbours == 3) {
				board[row][col] = 3
			} else if board[row][col] == 0 && liveNeighbours == 3 {
				board[row][col] = 2
			}
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			board[row][col] >>= 1
		}
	}
}
