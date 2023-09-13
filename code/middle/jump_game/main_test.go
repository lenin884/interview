package jump_game

import "testing"

func TestName(t *testing.T) {
	t.Run("can jump", func(t *testing.T) {
		nums := []int{2, 3, 1, 1, 4}

		got := canJump(nums)

		if !got {
			t.Errorf("got %v want %v", got, false)
		}
	})

	t.Run("not jump", func(t *testing.T) {
		nums := []int{3, 2, 1, 0, 4}

		got := canJump(nums)

		if got {
			t.Errorf("got %v want %v", got, true)
		}
	})
}
