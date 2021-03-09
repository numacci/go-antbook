package datastruct

import "fmt"

// PairQueue implementation
type PairQueue struct {
	data []*Pair
	size int
}

func NewPairQueue(cap int) *PairQueue {
	return &PairQueue{data: make([]*Pair, 0, cap), size: 0}
}

func (q *PairQueue) Push(p *Pair) {
	q.data = append(q.data, p)
	q.size++
}

func (q *PairQueue) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data[0] = nil
	q.data = q.data[1:]
	return true
}

func (q *PairQueue) Front() *Pair {
	return q.data[0]
}

func (q *PairQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *PairQueue) String() string {
	return fmt.Sprint(q.data)
}
