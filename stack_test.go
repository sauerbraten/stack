package stack

import (
	"fmt"
	"testing"
)

func printStack(s *Stack) {
	fmt.Println("- begin stack -")

	for i := range s.slice {
		fmt.Printf("\t%v\n", s.slice[len(s.slice)-1-i])
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

	if s.Pop() != nil {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	s.Push(4)
	s.Push("5")
	s.Push(6)

	if s.Pop().(int) != 6 {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	if s.Pop().(string) != "5" {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	if s.Pop().(int) != 4 {
		t.Log("Push/Pop did not work!")
		t.Fail()
	}

	if s.Pop() != nil {
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

	fmt.Printf("\n* popped element: %v *\n\n", s.Pop())
	printStack(s)

	fmt.Printf("\n* popped element: %v *\n\n", s.Pop())
	printStack(s)

	fmt.Printf("\n* popped element: %v *\n\n", s.Pop())
	printStack(s)

	fmt.Printf("\n* popped element: %v *\n\n", s.Pop())
	printStack(s)

	// Output:
	// - begin stack -
	// 	6
	// 	5
	// 	4
	// -  end stack  -
	//
	// * popped element: 6 *
	// 
	// - begin stack -
	// 	5
	// 	4
	// -  end stack  -
	//
	// * popped element: 5 *
	//
	// - begin stack -
	// 	4
	// -  end stack  -
	//
	// * popped element: 4 *
	//
	// - begin stack -
	// -  end stack  -
	//
	// * popped element: <nil> *
	//
	// - begin stack -
	// -  end stack  -
}
