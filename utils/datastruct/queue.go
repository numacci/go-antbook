package datastruct

import "fmt"

// Queue implementation
type Queue struct {
	data []int
	size int
}

func NewQueue(cap int) *Queue {
	return &Queue{data: make([]int, 0, cap), size: 0}
}

func (q *Queue) Push(n int) {
	q.data = append(q.data, n)
	q.size++
}

func (q *Queue) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

func (q *Queue) Front() int {
	return q.data[0]
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) String() string {
	return fmt.Sprint(q.data)
}
