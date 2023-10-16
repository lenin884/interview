package factorial_tralling_zeros

func trailingZeroes(n int) int {
	var result int

	for n > 0 {
		n /= 5
		result += n
	}

	return result
}
