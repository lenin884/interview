package min_stack

type MinStack struct {
	stack     []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minValues) == 0 {
		s.minValues = append(s.minValues, val)
		return
	}
	s.minValues = append(s.minValues, min(s.minValues[len(s.minValues)-1], val))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minValues = s.minValues[:len(s.minValues)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.minValues) == 0 {
		return 0
	}
	return s.minValues[len(s.minValues)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
