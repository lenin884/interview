package group_anagram

// time limit exceeded
func groupAnagrams(words []string) [][]string {
	result := make([][]string, 0)
	handleIndex := make(map[int]bool)

	for i := 0; i < len(words); i++ {
		if handleIndex[i] {
			continue
		}
		hashLetter := make(map[string]int)
		for _, letter := range words[i] {
			hashLetter[string(letter)] += 1
		}

		resultAnagram := make([]string, 0)
		resultAnagram = append(resultAnagram, words[i])
		for j := i + 1; j < len(words); j++ {
			if handleIndex[j] {
				continue
			}
			hashLetterCheck := make(map[string]int)
			for _, letter := range words[j] {
				hashLetterCheck[string(letter)] += 1
			}

			isAnagram := true
			for letterCheck, countCheck := range hashLetterCheck {
				if count, ok := hashLetter[letterCheck]; !ok || count != countCheck {
					isAnagram = false
				}
			}
			if isAnagram {
				resultAnagram = append(resultAnagram, words[j])
				handleIndex[j] = true
			}
		}
		handleIndex[i] = true
		result = append(result, resultAnagram)
	}

	return result
}

/*
best solution
*/

type Key [26]byte

func strKey(str string) (key Key) {
	for _, v := range str {
		key[v-'a']++
	}
	return
}

func groupAnagramsBest(strs []string) [][]string {
	groups := make(map[Key][]string)

	for _, str := range strs {
		key := strKey(str)
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for key := range groups {
		result = append(result, groups[key])
	}

	return result
}
