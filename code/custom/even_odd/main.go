package even_odd

func isEvenBinary(n int) bool {
	if 1&n == 1 {
		return false
	}
	return true
}

// isEven2NSolution checks if n is even using binary search between 2^0 and 2^32 using only plus and minus operations
func isEven2NSolution(n int) bool {
	if n == 0 {
		return true
	}
	if n == 1 {
		return false
	}
	leftValue := 0

	return searchBy2N(leftValue, n)
}

func searchBy2N(left, n int) bool {
	step := 1
	for {
		// находим интервал между степенями двойки
		if left+1<<step < n && n < left+1<<(step+1) {
			left = left + 1<<step
		} else {
			step++
			continue
		}

		if left == n {
			return true
		}

		if left+1 == n {
			return false
		}

		// продолжаем поиск
		return searchBy2N(left, n)
	}
}
