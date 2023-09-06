package word_pattern

import "testing"

func TestName(t *testing.T) {
	expected := []bool{false, false, true, false, false}
	patterns := []string{"aaa", "abba", "abba", "abba", "aaaa"}
	strings := []string{"aa aa aa aa", "dog dog dog dog", "dog cat cat dog", "dog cat cat fish", "dog cat cat dog"}

	for i := 0; i < len(expected); i++ {
		if wordPattern(patterns[i], strings[i]) != expected[i] {
			t.Errorf("\ninput pattern: %s\ninput s: %s\nexpected: %t\nactual: %t", patterns[i], strings[i], expected[i], !expected[i])
		}
	}
}
