package datastruct

import (
	"fmt"
	"testing"
)

func TestPairQueue(t *testing.T) {
	q := NewPairQueue(100)
	fmt.Println(q)

	if !q.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", q.IsEmpty())
	}

	q.Push(&Pair{1, 2})
	fmt.Println(q)
	q.Push(&Pair{-3, 1})
	fmt.Println(q)
	q.Push(&Pair{-1, -4})
	fmt.Println(q)

	if q.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", q.IsEmpty())
	}

	p1 := q.Front()
	fmt.Println(q)
	if p1.First != 1 || p1.Second != 2 {
		t.Fatalf("%v is expected, but %v\n", &Pair{1, 2}, p1)
	}
	q.Pop()

	p2 := q.Front()
	fmt.Println(q)
	if p2.First != -3 || p2.Second != 1 {
		t.Fatalf("%v is expected, but %v\n", &Pair{-3, 1}, p2)
	}
	q.Pop()

	p3 := q.Front()
	fmt.Println(q)
	if p3.First != -1 || p3.Second != -4 {
		t.Fatalf("%v is expected, but %v\n", &Pair{-1, -4}, p3)
	}
	q.Pop()

	fmt.Println(q)
	if !q.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", q.IsEmpty())
	}
}
