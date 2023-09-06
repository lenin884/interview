package word_pattern

import "strings"

/*
wordPattern("abba", "dog cat cat dog") -> true
need to replace variables names usedWords and repeatWords as more meaningful names
*/
func wordPattern(pattern string, s string) bool {
	usedWords := make(map[rune]string)
	repeatWords := make(map[string]rune)

	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	/*
		case when pattern = abba and s = dog dog dog dog
		a -> dog : usedWords = {a: dog}
		b -> dog : usedWords = {a: dog, b: dog} -> return false because a -> dog is already used
		b -> dog : usedWords = {a: dog, b: dog}
		a -> dog : usedWords = {a: dog, b: dog}
	*/
	for i, p := range pattern {
		if _, ok := usedWords[p]; ok {
			if usedWords[p] != words[i] {
				return false
			}

			continue
		}

		if _, ok := repeatWords[words[i]]; ok {
			if repeatWords[words[i]] != p {
				return false
			}
		}

		usedWords[p] = words[i]
		repeatWords[words[i]] = p
	}

	return true
}
