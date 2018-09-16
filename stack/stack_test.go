package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	if actual := stack.Empty(); actual != true {
		t.Errorf("Got %v expected %v", actual, true)
	}

	//insertion
	stack.Push(12)
	stack.Push(13)
	stack.Push(14)

	if actual := stack.Empty(); actual != false {
		t.Errorf("Got %v expected %v", actual, false)
	}

	if actual := stack.Size(); actual != 3 {
		t.Errorf("Got %v expected %v", actual, 3)
	}

	if actual, ok := stack.Peek(); actual != 14 || !ok {
		t.Errorf("Got %v expected %v", actual, 14)
	}

	stack.Pop()

	if actual, ok := stack.Peek(); actual != 13 || !ok {
		t.Errorf("Got %v expected %v", actual, 13)
	}

	if actual, ok := stack.Pop(); actual != 13 || !ok {
		t.Errorf("Got %v expected %v", actual, 13)
	}

	if actual, ok := stack.Pop(); actual != 12 || !ok {
		t.Errorf("Got %v expected %v", actual, 12)
	}

	if actual, ok := stack.Pop(); actual != nil || ok {
		t.Errorf("Got %v expected %v", actual, nil)
	}

	if actual := stack.Empty(); actual != true {
		t.Errorf("Got %v expected %v", actual, true)
	}
}
