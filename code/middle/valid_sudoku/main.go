package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	// 9 rows
	// 9 columns
	// 9 3x3 sub-boxes

	// 1. check rows
	for i := 0; i < 9; i++ {
		if !isValidRow(board, i) {
			return false
		}
	}

	// 2. check columns
	for i := 0; i < 9; i++ {
		if !isValidColumn(board, i) {
			return false
		}
	}

	// 3. check sub-boxes
	for i := 0; i < 9; i++ {
		if !isValidSubBox(board, i) {
			return false
		}
	}

	return true
}

func isValidRow(board [][]byte, row int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[row][i] == '.' {
			continue
		}
		if _, ok := m[board[row][i]]; ok {
			return false
		}
		m[board[row][i]] = true
	}
	return true
}

func isValidColumn(board [][]byte, column int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[i][column] == '.' {
			continue
		}
		if _, ok := m[board[i][column]]; ok {
			return false
		}
		m[board[i][column]] = true
	}
	return true
}

func isValidSubBox(board [][]byte, subBox int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		row := (subBox/3)*3 + i/3
		column := (subBox%3)*3 + i%3
		if board[row][column] == '.' {
			continue
		}
		if _, ok := m[board[row][column]]; ok {
			return false
		}
		m[board[row][column]] = true
	}
	return true
}
