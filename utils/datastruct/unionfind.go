package datastruct

import "fmt"

type UnionFind struct {
	par  []int
	rank []int
}

// NewUnionFind initialize UnionFind
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		par:  make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.par[i] = i
	}
	return uf
}

// Find returns the root of the tree including x
func (uf *UnionFind) Find(x int) int {
	if uf.par[x] == x {
		return x
	} else {
		uf.par[x] = uf.Find(uf.par[x])
		return uf.par[x]
	}
}

// Unite unites trees including x and y
func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.rank[x] < uf.rank[y] {
		uf.par[x] = y
	} else {
		uf.par[y] = x
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

// Same returns whether x and y belong to the same set
func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) String() string {
	return fmt.Sprintf("par = %v, rank = %v", uf.par, uf.rank)
}
