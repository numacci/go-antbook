package datastruct

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(100)

	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}

	stack.Push(10)
	stack.Push(1)
	stack.Push(-5)

	if stack.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", stack.IsEmpty())
	}

	if stack.Top() != -5 {
		t.Fatalf("-5 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if stack.Top() != 1 {
		t.Fatalf("1 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if stack.Top() != 10 {
		t.Fatalf("10 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}
}
