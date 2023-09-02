package is_subsequence

func isSubsequence(s, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	i := 0
	for j := 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}

		if i == len(s) {
			return true
		}
	}

	return false
}
