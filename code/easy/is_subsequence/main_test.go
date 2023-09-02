package is_subsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	expected := []bool{true, false, true, false}
	s := []string{"abc", "axc", "", "acb"}
	p := []string{"ahbgdc", "ahbgdc", "ahbgdc", "ahbgdc"}

	for i := 0; i < len(expected); i++ {
		if isSubsequence(s[i], p[i]) != expected[i] {
			t.Errorf("\ninput s: %s\ninput t: %s\nexpected: %t\nactual: %t", s[i], p[i], expected[i], !expected[i])
		}
	}
}
