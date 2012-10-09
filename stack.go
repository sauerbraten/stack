package stack

import (
	"errors"
)

// A stack.
type Stack []interface{}

// New returns an initialized stack.
func New() *Stack {
	return new(Stack)
}

// Push puts a new element on top of the stack.
func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

// Pop removes the top element and returns it, along with an error that is only != nil if the stack is empty.
func (s *Stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("stack is empty")
	}

	temp := (*s)[len(*s)-1]

	*s = (*s)[0 : len(*s)-1]

	return temp, nil
}

// Size returns the amount of elements in the stack.
func (s *Stack) Size() int {
	return len(*s)
}
