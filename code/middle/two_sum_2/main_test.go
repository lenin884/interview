package two_sum_2

import "testing"

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		input := []int{2, 7, 11, 15}
		target := 9
		expected := []int{1, 2}

		actual := twoSum(input, target)

		if len(actual) != len(expected) {
			t.Fatalf("expected %v, got %v", expected, actual)
		}

		for i := range actual {
			if actual[i] != expected[i] {
				t.Fatalf("expected %v, got %v", expected, actual)
			}
		}
	})
}
