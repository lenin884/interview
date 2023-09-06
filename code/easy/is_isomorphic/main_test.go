package is_isomorphic

import "testing"

func TestFunctional(t *testing.T) {
	expected := []bool{false, true, false, true, false}
	s := []string{"badc", "egg", "foo", "paper", "ab"}
	tt := []string{"baba", "add", "bar", "title", "aa"}

	for i := 0; i < len(expected); i++ {
		if isIsomorphic(s[i], tt[i]) != expected[i] {
			t.Errorf("\ninput s: %s\ninput t: %s\nexpected: %t\nactual: %t", s[i], tt[i], expected[i], !expected[i])
		}
	}
}
