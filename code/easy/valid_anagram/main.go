package valid_anagram

func isAnagram(s string, t string) bool {
	// check if length of s and t are equal
	if len(s) != len(t) {
		return false
	}

	// create a map to store the count of each character in s
	chars := make(map[rune]int, 0)

	for _, v := range s {
		chars[v]++
	}

	// iterate over t and decrement the count of each character in chars
	for _, v := range t {
		chars[v]--
	}

	// iterate over chars and check if the count of each character is 0
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
