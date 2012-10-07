package stack

// An element.
type element struct {
	next  *element
	value interface{}
}

// A stack.
type Stack struct {
	top *element
}

// New returns an initialized stack.
func New() *Stack {
	return &Stack{nil}
}

// Push puts a new element on top of the stack.
func (s *Stack) Push(value interface{}) {
	temp := s.top
	newE := &element{temp, value}
	s.top = newE
}

// Pop removes the top element and returns it.
func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	v := s.top.value
	s.top = s.top.next
	return v
}

// Size returns the amount of elements in the stack.
func (s *Stack) Size() int {
	i := 0

	// if stack empty
	if s.top == nil {
		return i
	}

	e := s.top
	i++

	for e.next != nil {
		e = e.next
		i++
	}

	return i
}
