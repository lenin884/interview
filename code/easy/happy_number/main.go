package happy_number

func isHappy(n int) bool {
	// create a map to store the sum of each digit in n
	sums := make(map[int]bool, 0)

	// iterate until n is 1
	for n != 1 {
		// calculate the sum of each digit in n
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}

		// check if the sum is in sums
		if sums[sum] {
			return false
		}

		// add the sum to sums
		sums[sum] = true

		// set n to sum
		n = sum
	}

	return true
}
