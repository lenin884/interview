package jump_game_2

import "testing"

func TestName(t *testing.T) {
	t.Run("5 elements got 2", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}

		got := jump(nums)
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
