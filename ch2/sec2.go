package ch2

import (
	"container/heap"
	"github.com/numacci/go-antbook/utils/datastruct"
	"github.com/numacci/go-antbook/utils/math"
	"sort"
)

func p042(c []int, A int) int {
	v := []int{1, 5, 10, 50, 100, 500}
	ans := 0
	for i := len(c) - 1; i >= 0; i-- {
		n := math.Min(c[i], A/v[i])
		A -= n * v[i]
		ans += n
	}
	if A != 0 {
		return -1
	}
	return ans
}

func p043(n int, s []int, t []int) int {
	p := make([][2]int, n)
	for i := 0; i < n; i++ {
		p[i][0] = t[i]
		p[i][1] = s[i]
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i][0] < p[j][0]
	})

	ans, et := 0, 0
	for i := 0; i < n; i++ {
		if et < p[i][1] {
			et = p[i][0]
			ans++
		}
	}
	return ans
}

func p045(N int, S string) string {
	reverse := func(s string) string {
		rs := []rune(s)
		for i := 0; i < N/2; i++ {
			rs[i], rs[N-1-i] = rs[N-1-i], rs[i]
		}
		return string(rs)
	}

	rS := reverse(S)
	T := ""
	for len(T) < N {
		if S < rS {
			T += S[:1]
			S = S[1:]
		} else {
			T += rS[:1]
			rS = rS[1:]
		}
	}
	return T
}

func p047(N, R int, X []int) int {
	i, ans := 0, 0
	for i < N {
		l := X[i]
		i++
		for i < N && X[i]-l <= R {
			i++
		}
		// painted point
		p := X[i-1]
		for i < N && X[i]-p <= R {
			i++
		}

		ans++
	}
	return ans
}

func p049(N int, L []int) int {
	ih := &datastruct.IntHeap{}
	heap.Init(ih)
	for _, n := range L {
		heap.Push(ih, n)
	}

	ans := 0
	for ih.Len() > 1 {
		l1 := heap.Pop(ih).(int)
		l2 := heap.Pop(ih).(int)
		heap.Push(ih, l1+l2)
		ans += l1 + l2
	}
	return ans
}
