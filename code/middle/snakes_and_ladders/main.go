package snakes_and_ladders

/*
Using BFS method
Need to draw for better understanding
*/
func snakesAndLadders(board [][]int) int {
	queue := []int{1}

	visited := make(map[int]bool)

	moves := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == len(board)*len(board) {
				return moves
			}

			if visited[node] {
				continue
			}

			visited[node] = true

			for j := 1; j <= 6; j++ {
				next := node + j

				if next > len(board)*len(board) {
					break
				}

				row, col := getRowCol(next, len(board))

				if board[row][col] != -1 {
					next = board[row][col]
				}

				queue = append(queue, next)
			}
		}

		moves++
	}

	return -1
}

func getRowCol(node int, n int) (int, int) {
	row := (node - 1) / n
	col := (node - 1) % n

	if row%2 == 1 {
		col = n - 1 - col
	}

	row = n - 1 - row

	return row, col
}
