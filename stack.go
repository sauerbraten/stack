package stack

// A stack.
type Stack struct {
	slice []interface{}
}

// New returns an initialized stack.
func New() *Stack {
	return new(Stack)
}

// Push puts a new element on top of the stack.
func (s *Stack) Push(value interface{}) {
	s.slice = append(s.slice, value)
}

// Pop removes the top element and returns it.
func (s *Stack) Pop() interface{} {
	if len(s.slice) == 0 {
		return nil
	}

	temp := s.slice[len(s.slice)-1]

	s.slice = s.slice[0 : len(s.slice)-1]

	return temp
}

// Size returns the amount of elements in the stack.
func (s *Stack) Size() int {
	return len(s.slice)
}
