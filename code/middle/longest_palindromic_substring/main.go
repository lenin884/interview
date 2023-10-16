package longest_palindromic_substring

// longestPalindrome is function to find longest palindromic substring with dynamic programming
func longestPalindrome(s string) string {
	// get length of string
	n := len(s)
	// create 2D array with size n x n
	// dp[i][j] is true if substring from i to j is palindromic
	dp := make([][]bool, n)

	// create variable to store start and end index of longest palindromic substring
	start, end := 0, 0

	// loop from n - 1 to 0
	for i := n - 1; i >= 0; i-- {
		// set dp[i][i] to true
		dp[i] = make([]bool, n)
		dp[i][i] = true

		// loop from i + 1 to n
		for j := i + 1; j < n; j++ {
			// if s[i] is equal to s[j] and dp[i + 1][j - 1] is true
			if s[i] == s[j] && (j-i == 1 || dp[i+1][j-1]) {
				// set dp[i][j] to true
				dp[i][j] = true
				// if length of substring from i to j is greater than length of substring from start to end
				if j-i > end-start {
					// set start to i and end to j
					start, end = i, j
				}
			}
		}
	}

	// return substring from start to end
	return s[start : end+1]
}

// longestPalindrome2 is function to find longest palindromic substring with expand around center
func longestPalindrome2(s string) string {
	// get length of string
	n := len(s)

	// create variable to store start and end index of longest palindromic substring
	start, end := 0, 0

	// loop from 0 to n
	for i := 0; i < n; i++ {
		// get length of palindromic substring with i as center
		len1 := expandAroundCenter(s, i, i)
		// get length of palindromic substring with i and i + 1 as center
		len2 := expandAroundCenter(s, i, i+1)
		// get max length of palindromic substring
		ln := max(len1, len2)
		// if max length is greater than end - start
		if ln > end-start {
			// set start to i - (len - 1) / 2
			start = i - (ln-1)/2
			// set end to i + len / 2
			end = i + ln/2
		}
	}

	// return substring from start to end
	return s[start : end+1]
}

// expandAroundCenter is function to get length of palindromic substring with center i and j
func expandAroundCenter(s string, i, j int) int {
	// get length of string
	n := len(s)

	// loop while i is greater than or equal to 0 and j is less than n and s[i] is equal to s[j]
	for i >= 0 && j < n && s[i] == s[j] {
		// decrement i
		i--
		// increment j
		j++
	}

	// return length of palindromic substring
	return j - i - 1
}
