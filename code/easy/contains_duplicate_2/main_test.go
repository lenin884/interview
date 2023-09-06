package contains_duplicate_2

import "testing"

func TestName(t *testing.T) {
	expected := []bool{true, true, false}
	inputs := [][]int{{1, 2, 3, 1}, {1, 0, 1, 1}, {1, 2, 3, 1, 2, 3}}
	ks := []int{3, 1, 2}
	for i, exp := range expected {
		if containsNearbyDuplicate(inputs[i], ks[i]) != exp {
			t.Errorf("\ninput: %v\nexpected: %t\nactual: %t", inputs[i], exp, !exp)
		}
	}
}
