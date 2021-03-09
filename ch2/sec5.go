package ch2

import (
	"container/heap"
	"fmt"
	"github.com/numacci/go-antbook/utils/datastruct"
	"github.com/numacci/go-antbook/utils/graph"
	"sort"
)

func dfsG(v, c int, edges [][]*graph.Edge, color []int) bool {
	color[v] = c
	for i := 0; i < len(edges[v]); i++ {
		e := edges[v][i]
		if color[e.To] == c {
			return false
		}
		if color[e.To] == 0 && !dfsG(e.To, -c, edges, color) {
			return false
		}
	}
	return true
}

func p093(n int, mat [][]int) string {
	edges := make([][]*graph.Edge, n)
	for i := 0; i < len(mat); i++ {
		s, t := mat[i][0], mat[i][1]
		edges[s] = append(edges[s], &graph.Edge{To: t})
		edges[t] = append(edges[t], &graph.Edge{To: s})
	}

	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			if !dfsG(i, 1, edges, color) {
				return "No"
			}
		}
	}
	return "Yes"
}

func kruskal(V, E int, es []*graph.Edge) int {
	// コストが小さい順にソート
	sort.Slice(es, func(i, j int) bool {
		return es[i].Cost < es[j].Cost
	})

	uf := datastruct.NewUnionFind(V)
	res := 0
	for _, e := range es {
		if !uf.Same(e.From, e.To) {
			uf.Unite(e.From, e.To)
			res += e.Cost
		}
	}
	return res
}

func p101(V, E int, es []*graph.Edge) int {
	return kruskal(V, E, es)
}

func p102(N, R int, gh [][]int) int {
	G := make([][]*graph.Edge, N)
	for i := 0; i < N; i++ {
		G[i] = make([]*graph.Edge, 0, N)
	}
	for i := 0; i < R; i++ {
		from, to, cost := gh[i][0]-1, gh[i][1]-1, gh[i][2]
		G[from] = append(G[from], &graph.Edge{To: to, Cost: cost})
		G[to] = append(G[to], &graph.Edge{To: from, Cost: cost})
	}

	dist1 := make([]int, N)
	dist2 := make([]int, N)
	for i := 0; i < N; i++ {
		dist1[i], dist2[i] = Inf, Inf
	}

	pq := datastruct.NewPriorityQueue(N * R)
	heap.Init(&pq)

	dist1[0] = 0
	heap.Push(&pq, &datastruct.Pair{First: 0, Second: 0})
	for !pq.IsEmpty() {
		p := heap.Pop(&pq).(*datastruct.Pair)
		d, v := p.First, p.Second
		if dist2[v] < d {
			continue
		}

		for _, e := range G[v] {
			d2 := d + e.Cost
			if dist1[e.To] > d2 {
				dist1[e.To], d2 = d2, dist1[e.To]
				heap.Push(&pq, &datastruct.Pair{First: dist1[e.To], Second: e.To})
			}
			if dist2[e.To] > d2 {
				dist2[e.To] = d2
				heap.Push(&pq, &datastruct.Pair{First: dist2[e.To], Second: e.To})
			}
		}
	}
	return dist2[N-1]
}

func p103(N, M, R int, rel [][]int) int {
	es := make([]*graph.Edge, R)
	for i := 0; i < R; i++ {
		es[i] = &graph.Edge{From: rel[i][0], To: N + rel[i][1], Cost: -rel[i][2]}
	}
	return (N+M)*10000 + kruskal(N+M, R, es)
}

func p104(N, ML, MD int, ok, ng [][]int) int {
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = Inf
	}

	// d[i] <= d[i+1] && d[BL] - d[AL] <= DL && d[BD] - d[AD] >= DD
	// i => i+1 へコスト0の辺
	// AL => BL へコストDLの辺
	// BD => AD へコスト-DDの辺
	V, E := N, N+ML+MD-1
	es := make([]*graph.Edge, 0, E)
	for i := 0; i < N-1; i++ {
		es = append(es, &graph.Edge{From: i + 1, To: i, Cost: 0})
	}
	for i := 0; i < ML; i++ {
		es = append(es, &graph.Edge{From: ok[i][0] - 1, To: ok[i][1] - 1, Cost: ok[i][2]})
	}
	for i := 0; i < MD; i++ {
		es = append(es, &graph.Edge{From: ng[i][1] - 1, To: ng[i][0] - 1, Cost: -ng[i][2]})
	}
	fmt.Println("es", es)

	gh := &graph.Graph{V: V, E: E}
	dist, neg := gh.BellmanFord(es, 0, true)
	res := dist[N-1]
	if neg {
		res = -1
	} else if dist[N-1] == graph.Inf {
		res = -2
	}
	fmt.Println(dist, ng)
	return res
}
