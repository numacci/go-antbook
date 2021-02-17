package datastruct

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	s1 := "string 1"
	s2 := s1
	fmt.Println(s1, s2)
	fmt.Println(&s1, &s2)

	s2 = "string 2"
	fmt.Println(s1, s2)
	fmt.Println(&s1, &s2)

	queue := NewQueue(10)

	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}

	queue.Push(10)
	queue.Push(1)
	queue.Push(-5)

	if queue.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", queue.IsEmpty())
	}

	if queue.Front() != 10 {
		t.Fatalf("10 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if queue.Front() != 1 {
		t.Fatalf("1 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if queue.Front() != -5 {
		t.Fatalf("-5 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}
}
