package valid_anagram

import "testing"

func TestName(t *testing.T) {
	expected := []bool{true, true, false, true}
	s := []string{"anagram", "art", "rat", "курва"}
	tt := []string{"nagaram", "rat", "car", "курав"}

	for i := 0; i < len(expected); i++ {
		if isAnagram(s[i], tt[i]) != expected[i] {
			t.Errorf("\ninput s: %s\ninput t: %s\nexpected: %t\nactual: %t", s[i], tt[i], expected[i], !expected[i])
		}
	}
}
