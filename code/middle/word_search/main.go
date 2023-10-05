package word_search

func exist(board [][]byte, word string) bool {
	for i, row := range board {
		for j, _ := range row {
			if search(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return false
	}
	if board[i][j] != word[0] {
		return false
	}
	board[i][j] = '#'
	if search(board, i+1, j, word[1:]) ||
		search(board, i-1, j, word[1:]) ||
		search(board, i, j+1, word[1:]) ||
		search(board, i, j-1, word[1:]) {
		return true
	}
	board[i][j] = word[0]
	return false
}
