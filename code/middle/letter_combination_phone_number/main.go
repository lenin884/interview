package letter_combination_phone_number

var letterMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)

	if len(digits) != 0 {
		backtrack(0, []byte{}, digits, &result)
	}

	return result
}

func backtrack(i int, combinations []byte, digits string, result *[]string) {
	if len(combinations) == len(digits) {
		*result = append(*result, string(combinations))
		return
	}

	for _, letter := range letterMap[digits[i]] {
		combinations = append(combinations, letter)
		backtrack(i+1, combinations, digits, result)
		combinations = combinations[:len(combinations)-1]
	}
}
