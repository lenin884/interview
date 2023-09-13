package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int

	for temp := x; temp > 0; temp /= 10 {
		reversed = reversed*10 + temp%10
	}

	return reversed == x
}
