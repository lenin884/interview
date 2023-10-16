package add_binary

import "strconv"

func addBinary(a string, b string) string {
	var result string

	ai := len(a) - 1
	bi := len(b) - 1
	carry := 0

	for ai >= 0 || bi >= 0 {
		// full adder input : carry, x, y
		x, y := 0, 0
		if ai >= 0 {
			x = int(a[ai] - '0')
			ai--
		}
		if bi >= 0 {
			y = int(b[bi] - '0')
			bi--
		}

		// full adder output : carry, sum
		sum := strconv.Itoa(carry ^ x ^ y)
		carry = (carry & x) | (x & y) | (carry & y)

		result = sum + result
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}
