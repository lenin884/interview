package happy_number

import "testing"

func TestName(t *testing.T) {
	expected := []bool{true, false}
	inputs := []int{19, 2}

	for i := 0; i < len(expected); i++ {
		if isHappy(inputs[i]) != expected[i] {
			t.Errorf("\ninput: %d\nexpected: %t\nactual: %t", inputs[i], expected[i], !expected[i])
		}
	}
}
