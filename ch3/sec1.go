package ch3

import (
	"github.com/numacci/go-antbook/utils/math"
	math2 "math"
	"sort"
)

func p128(n, k int, a []int) int {
	return math.LowerBound(a, k)
}

func c129(K int, l float64, L []float64) bool {
	n := 0
	for _, li := range L {
		n += int(li / l)
	}
	return n >= K
}

func p129(N, K int, L []float64) float64 {
	sort.Float64s(L)

	l, r := 0.0, 1e6
	for i := 0; i < 100; i++ {
		mid := (l + r) / 2

		if c129(K, mid, L) {
			l = mid
		} else {
			r = mid
		}
	}
	return math2.Floor(r*100) / 100
}

func c131(N, M, d int, x []int) bool {
	last := 0
	for i := 1; i < M; i++ {
		cur := last + 1
		for cur < N && x[cur]-x[last] < d {
			cur++
		}
		if cur == N {
			return false
		}
		last = cur
	}
	return true
}

func p131(N, M int, x []int) int {
	sort.Ints(x)

	l, r := 0, int(1e9)
	for l+1 < r {
		mid := (l + r) / 2
		if c131(N, M, mid, x) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

func c132(n, k int, w, v []int, x float64) bool {
	// sum(v)/sum(w) >= x を満たす品物の集合が存在するか
	// <=> sum(vi - x*wi) >= 0
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		y[i] = float64(v[i]) - x*float64(w[i])
	}

	sum := 0.0
	sort.Float64s(y)
	for i := 0; i < k; i++ {
		sum += y[n-i-1]
	}
	return sum >= 0
}

func p132(n, k int, w, v []int) float64 {
	l, r := 0.0, 1e9
	for i := 0; i < 100; i++ {
		mid := (l + r) / 2
		if c132(n, k, w, v, mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return math2.Floor(r*1e6) / 1e6
}
