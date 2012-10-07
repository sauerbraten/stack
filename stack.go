package stack

// An element.
type element struct {
	next  *element
	value interface{}
}

// A stack.
type Stack struct {
	top  *element
	size int
}

// New returns an initialized stack.
func New() *Stack {
	return &Stack{nil, 0}
}

// Push puts a new element on top of the stack.
func (s *Stack) Push(value interface{}) {
	temp := s.top
	newE := &element{temp, value}
	s.top = newE
	size++
}

// Pop removes the top element and returns it.
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	v := s.top.value
	s.top = s.top.next
	size--
	return v
}

// Size returns the amount of elements in the stack.
func (s *Stack) Size() int {
	return s.size
}
