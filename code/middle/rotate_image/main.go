package rotate_image

func rotate(matrix [][]int) {
	left, right := 0, len(matrix)-1

	for left < right {
		for i := 0; i < right-left; i++ {
			top, bottom := left, right

			// save top left
			topLeft := matrix[top][left+i]

			// move bottom left to top left
			matrix[top][left+i] = matrix[bottom-i][left]

			// move bottom right to bottom left
			matrix[bottom-i][left] = matrix[bottom][right-i]

			// move top right to bottom right
			matrix[bottom][right-i] = matrix[top+i][right]

			// move top left to top right
			matrix[top+i][right] = topLeft
		}

		left++
		right--
	}
	return
}

func rotateTranspose(matrix [][]int) {
	// Step 1: Transpose the matrix
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row of the transposed matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[0])-j-1] = matrix[i][len(matrix[0])-j-1], matrix[i][j]
		}
	}
}
