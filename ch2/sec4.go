package ch2

import (
	"container/heap"
	"github.com/numacci/go-antbook/utils/datastruct"
)

func p073(N, L, P int, A, B []int) int {
	ih := &datastruct.IntHeap{}
	heap.Init(ih)

	A = append(A, L)
	B = append(B, 0)

	ans, pos := 0, 0
	for i := 0; i <= N; i++ {
		P -= A[i] - pos
		for P < 0 {
			if ih.Len() == 0 {
				return -1
			}
			P += -heap.Pop(ih).(int)
			ans++
		}
		pos = A[i]
		heap.Push(ih, -B[i])
	}
	return ans
}

func p085(N, K int, T, X, Y []int) int {
	ans := 0
	uf := datastruct.NewUnionFind(3 * N)
	for i := 0; i < K; i++ {
		// invalid number
		if X[i] < 1 || X[i] > 100 || Y[i] < 1 || Y[i] > 100 {
			ans++
			continue
		}

		t, x, y := T[i], X[i]-1, Y[i]-1
		if t == 1 {
			// xがyを食べる or xがyに食べられる条件が既に適用されている場合
			if uf.Same(x, y+N) || uf.Same(x, y+2*N) {
				ans++
			} else {
				uf.Unite(x, y)
				uf.Unite(x+N, y+N)
				uf.Unite(x+2*N, y+2*N)
			}
		} else {
			// xとyが同じ種類 or xがyに食べられる条件が既に適用されている場合
			if uf.Same(x, y) || uf.Same(x, y+2*N) {
				ans++
			} else {
				uf.Unite(x, y+N)
				uf.Unite(x+N, y+2*N)
				uf.Unite(x+2*N, y)
			}
		}
	}
	return ans
}
