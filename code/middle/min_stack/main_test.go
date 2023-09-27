package min_stack

import "testing"

func TestName(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		if minStack.GetMin() != -3 {
			t.Errorf("GetMin() got %d, want %d", minStack.GetMin(), -3)
		}
		minStack.Pop()
		if minStack.Top() != 0 {
			t.Errorf("Top() got %d, want %d", minStack.Top(), 0)
		}
		if minStack.GetMin() != -2 {
			t.Errorf("GetMin() got %d, want %d", minStack.GetMin(), -2)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		minStack := Constructor()
		minStack.Push(0)
		minStack.Push(1)
		minStack.Push(0)
		if minStack.GetMin() != 0 {
			t.Errorf("GetMin() got %d, want %d", minStack.GetMin(), 0)
		}
		minStack.Pop()
		if minStack.GetMin() != 0 {
			t.Errorf("GetMin() got %d, want %d", minStack.GetMin(), 0)
		}
	})
}
