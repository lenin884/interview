package evaluate_reverse_polish_notation

import "testing"

func TestName(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		got := evalRPN([]string{"2", "1", "+", "3", "*"})
		want := 9
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		got := evalRPN([]string{"4", "13", "5", "/", "+"})
		want := 6
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
