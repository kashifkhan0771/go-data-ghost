package stack

type Stack struct {
	items []int
}

// NewStack creates and returns a new stack
func NewStack() *Stack {
	return &Stack{}
}

// Push pushes a value onto the top of the stack
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// Pop pops the top value off the stack
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, true
}

// Peek returns the top value on the stack
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// Size returns the size of the stack
func (s *Stack) Size() int {
	return len(s.items)
}
