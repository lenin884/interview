package reverse_words_in_string

import "strings"

func reverseWords(s string) string {
	wordlist := strings.Fields(s)

	for i := 0; i < len(wordlist)/2; i++ {
		wordlist[i], wordlist[len(wordlist)-i-1] = wordlist[len(wordlist)-i-1], wordlist[i]
	}

	return strings.TrimSpace(strings.Join(wordlist, " "))
}
