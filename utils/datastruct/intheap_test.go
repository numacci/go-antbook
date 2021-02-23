package datastruct

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println(*h)
	if (*h)[0] != 1 {
		t.Fatalf("1 is expected, but %v\n", (*h)[0])
	}

	i := 0
	expected := []int{1, 2, 3, 5}
	for h.Len() > 0 {
		if v := heap.Pop(h); v != expected[i] {
			t.Fatalf("%v is expected, but %v\n", expected[i], v)
		}
		i++
	}
}
