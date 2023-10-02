package surrounded_regions

func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	rows, cols := len(board), len(board[0])

	if i < 0 || i == rows || j < 0 || j == cols || board[i][j] != 'O' {
		return
	}

	board[i][j] = 'A'

	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)

	return
}
