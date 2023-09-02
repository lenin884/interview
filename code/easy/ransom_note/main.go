package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 {
		return true
	}

	if len(magazine) == 0 {
		return false
	}

	chars := make(map[rune]int)
	for _, r := range magazine {
		chars[r]++
	}

	for _, m := range ransomNote {
		if _, ok := chars[m]; !ok {
			return false
		}

		chars[m]--

		if chars[m] < 0 {
			return false
		}
	}

	return true
}
