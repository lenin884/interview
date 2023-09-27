package set_matrix_zeros

// setZeroesHashMap O(mn) time and O(m+n) space
func setZeroesHashMap(matrix [][]int) {
	mRow := make(map[int]bool)
	mCol := make(map[int]bool)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				mRow[i] = true
				mCol[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if _, ok := mRow[i]; ok {
				matrix[i][j] = 0
			}

			if _, ok := mCol[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}

// setZeroes O(mn) time and O(1) space
func setZeroes(matrix [][]int) {
	// Step 1: Mark the first row and column if they have zeroZ0[H
	firstRowHasZero := false
	firstColHasZero := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// Step 2: Use first row and column to mark zero
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0 // mark column
				matrix[i][0] = 0 // mark row
			}
		}
	}

	// Step 3: Set zeros except first row and column
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0 // set zero
			}
		}
	}

	// Step 4: Set zeros for first row and column mif needed
	if firstRowHasZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if firstColHasZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	return
}
