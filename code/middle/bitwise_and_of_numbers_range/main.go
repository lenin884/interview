package bitwise_and_of_numbers_range

func rangeBitwiseAnd(left int, right int) int {
	shift := 0

	// Находим количество сдвигов, необходимых для того, чтобы сравнить left и right
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}

	// Производим сдвиг обратно, чтобы получить общий префикс
	return left << shift
}
