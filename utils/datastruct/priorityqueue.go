package datastruct

import "fmt"

// PriorityQueue implementation
type PriorityQueue []*Item

type Item struct {
	P     *Pair
	Index int
}

func (i *Item) String() string {
	return fmt.Sprint(i.P)
}

func NewPriorityQueue(cap int) PriorityQueue {
	return make(PriorityQueue, 0, cap)
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].P.First < pq[j].P.First }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq PriorityQueue) IsEmpty() bool {
	return pq.Len() <= 0
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	p := x.(*Pair)
	i := &Item{
		P:     p,
		Index: n,
	}
	*pq = append(*pq, i)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	old[n-1] = nil
	i.Index = -1
	*pq = old[0 : n-1]
	return i.P
}
