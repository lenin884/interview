package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	hashmap := make(map[byte]int)

	if len(s) <= 1 {
		return len(s)
	}

	result := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if v, ok := hashmap[s[right]]; ok && v >= left {
			left = v + 1
		}

		hashmap[s[right]] = right
		result = max(result, right-left+1)
	}

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
