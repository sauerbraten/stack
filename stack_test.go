package stack

import (
	"fmt"
	"testing"
)

func printStack(s *Stack) {
	fmt.Println("- begin stack -")

	for i := range *s {
		fmt.Printf("\t%v\n", (*s)[len(*s)-1-i])
	}

	fmt.Println("-  end stack  -")
}

func TestSize(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Log("Size is wrong!")
		t.Fail()
	}

	s.Pop()
	if s.Size() != 2 {
		t.Log("Size is wrong!")
		t.Fail()
	}

	s.Pop()
	if s.Size() != 1 {
		t.Log("Size is wrong!")
		t.Fail()
	}

	s.Pop()
	if s.Size() != 0 {
		t.Log("Size is wrong!")
		t.Fail()
	}
}

func TestStack(t *testing.T) {
	s := New()

	if val, err := s.Pop(); val != nil || err == nil {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	s.Push(4)
	s.Push("5")
	s.Push(6)

	if val, err := s.Pop(); val.(int) != 6 || err != nil {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	if val, err := s.Pop(); val.(string) != "5" || err != nil {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	if val, err := s.Pop(); val.(int) != 4 || err != nil {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	if val, err := s.Pop(); val != nil || err == nil {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}
}

func ExampleStack() {
	s := New()

	s.Push(4)
	s.Push("5")
	s.Push(6)

	printStack(s)

	val, err := s.Pop()
	fmt.Printf("\n* popped element: %v with error: %v*\n\n", val, err)
	printStack(s)

	val, err = s.Pop()
	fmt.Printf("\n* popped element: %v with error: %v*\n\n", val, err)
	printStack(s)

	val, err = s.Pop()
	fmt.Printf("\n* popped element: %v with error: %v*\n\n", val, err)
	printStack(s)

	val, err = s.Pop()
	fmt.Printf("\n* popped element: %v with error: %v*\n\n", val, err)
	printStack(s)

	// Output:
	// - begin stack -
	// 	6
	// 	5
	// 	4
	// -  end stack  -
	//
	// * popped element: 6 with error: <nil>*
	// 
	// - begin stack -
	// 	5
	// 	4
	// -  end stack  -
	//
	// * popped element: 5 with error: <nil>*
	//
	// - begin stack -
	// 	4
	// -  end stack  -
	//
	// * popped element: 4 with error: <nil>*
	//
	// - begin stack -
	// -  end stack  -
	//
	// * popped element: <nil> with error: stack is empty*
	//
	// - begin stack -
	// -  end stack  -
}
