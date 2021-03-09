package graph

import (
	"container/heap"
	"fmt"
	"github.com/numacci/go-antbook/utils/datastruct"
	"github.com/numacci/go-antbook/utils/math"
)

const (
	Inf = int(1e9)
)

type Edge struct {
	From, To, Cost int
}

func (e *Edge) String() string {
	return fmt.Sprintf("[from=%v, to=%v, cost=%v]", e.From, e.To, e.Cost)
}

type Graph struct {
	V, E int
	Prev []int
}

// BellmanFord returns the shortest path array and the existence of negative loop
// Use BellmanFord only if some edges have the negative cost
func (g *Graph) BellmanFord(es []*Edge, s int, directed bool) ([]int, bool) {
	d := make([]int, g.V) // shortest path
	for i := 0; i < g.V; i++ {
		d[i] = Inf
	}

	E := g.E
	if !directed {
		E *= 2
		revEs := make([]*Edge, g.E)
		for i := 0; i < g.E; i++ {
			revEs[i] = &Edge{From: es[i].To, To: es[i].From, Cost: es[i].Cost}
		}
		es = append(es, revEs...)
	}

	d[s] = 0
	hasNegativeLoop := false
	for i := 0; i < g.V; i++ {
		updated := false
		for _, e := range es {
			if d[e.To] > d[e.From]+e.Cost {
				d[e.To] = d[e.From] + e.Cost
				updated = true

				if i == g.V-1 {
					hasNegativeLoop = true
				}
			}
		}
		if !updated {
			break
		}
	}
	return d, hasNegativeLoop
}

// Dijkstra returns the length of the shortest path and shortest path array
// Do not use Dijkstra when negative edges exist
func (g *Graph) Dijkstra(G [][]*Edge, s int) []int {
	d := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		d[i] = Inf
	}

	// PriorityQueue: (shortest path from the start vertex, index of vertex)
	pq := make(datastruct.PriorityQueue, 0, g.V*g.E)
	heap.Init(&pq)

	d[s] = 0
	heap.Push(&pq, &datastruct.Pair{First: 0, Second: s})

	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*datastruct.Pair)
		dist, v := p.First, p.Second

		if d[v] < dist {
			continue
		}

		for _, e := range G[v] {
			if d[e.To] > d[v]+e.Cost {
				d[e.To] = d[v] + e.Cost
				heap.Push(&pq, &datastruct.Pair{First: d[e.To], Second: e.To})
				// 最短経路まで求める必要がある場合
				if g.Prev != nil {
					g.Prev[e.To] = v

				}
			}
		}
	}
	return d
}

// FloydWarshall returns the shortest path matrix
func (g *Graph) FloydWarshall(G [][]int) [][]int {
	d := make([][]int, g.V)
	copy(d, G)

	for k := 0; k < g.V; k++ {
		for i := 0; i < g.V; i++ {
			for j := 0; j < g.V; j++ {
				d[i][j] = math.Min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	return d
}

// GetPath is only implemented for Dijkstra
func (g *Graph) GetPath(to int) []int {
	path := make([]int, 0, g.V)
	cur := to
	for ; cur != -1; cur = g.Prev[cur] {
		path = append(path, cur)
	}
	// Reverse
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}
	return path
}
