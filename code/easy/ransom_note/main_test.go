package ransom_note

import "testing"

func TestRansomNote(t *testing.T) {
	expected := []bool{true, true, true, false}
	ransomNotes := []string{"a", "b", "", "acx"}
	magazines := []string{"ab", "ba", "ab", "abc"}
	for i := 0; i < len(expected); i++ {
		if canConstruct(ransomNotes[i], magazines[i]) != expected[i] {
			t.Errorf("\ninput s: %s\ninput t: %s\nexpected: %t\nactual: %t", ransomNotes[i], magazines[i], expected[i], !expected[i])
		}
	}
}
