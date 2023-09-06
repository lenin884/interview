package is_isomorphic

func isIsomorphic(s string, t string) bool {
	chars := make(map[rune]rune, 0)
	used := make(map[rune]bool, 0)

	// check case when s = badc and t = baba
	// b -> b : used = {b: true}
	// a -> a : used = {b: true, a: true}
	// d -> b : used = {b: true, a: true, b: true}
	// c -> a
	for i, v := range s {
		if _, ok := chars[v]; !ok {
			if used[rune(t[i])] {
				return false
			}

			chars[v] = rune(t[i])
			used[rune(t[i])] = true
		}

		if chars[v] != rune(t[i]) {
			return false
		}
	}

	return true
}
