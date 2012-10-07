package stack

type Element struct {
	next  *Element
	value interface{}
}

type Stack struct {
	top *Element
}

func New() *Stack {
	return &Stack{nil}
}

func (s *Stack) Push(value interface{}) {
	temp := s.top
	newE := &Element{temp, value}
	s.top = newE
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}

	v := s.top.value
	s.top = s.top.next
	return v
}

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
