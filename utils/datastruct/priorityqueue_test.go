package datastruct

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	//pq := make(PriorityQueue, 0, 1024)
	pq := NewPriorityQueue(1024)
	heap.Init(&pq)
	heap.Push(&pq, &Pair{First: 3, Second: 1})
	heap.Push(&pq, &Pair{First: 10, Second: -3})
	heap.Push(&pq, &Pair{First: 1, Second: 4})
	fmt.Println(pq)

	if pq[0].P.First != 1 {
		t.Fatalf("1 is expected, but %v\n", pq[0].P.First)
	}

	i := 0
	expected := []*Pair{
		{1, 4},
		{3, 1},
		{10, -3},
	}
	for !pq.IsEmpty() {
		v := heap.Pop(&pq).(*Pair)
		if *v != *expected[i] {
			t.Fatalf("%v is expected, but %v\n", expected[i], v)
		}
		fmt.Println(pq)
		i++
	}
}
