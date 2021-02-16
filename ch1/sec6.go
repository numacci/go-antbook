package ch1

import (
	"math"
	"sort"
)

func p021(n int, a []int) int {
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				l := a[i] + a[j] + a[k]
				m := max(a[i], max(a[j], a[k]))

				if l-m > m {
					ans = max(ans, l)
				}
			}
		}
	}
	return ans
}

func p023(L int, n int, x []int) (int, int) {
	minT := 0
	for _, xi := range x {
		minT = max(minT, min(xi, L-xi))
	}
	maxT := 0
	for _, xi := range x {
		maxT = max(maxT, max(xi, L-xi))
	}
	return minT, maxT
}

func p027(n int, m int, k []int) string {
	kk := make([]int, n*n, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			kk[i*n+j] = k[i] + k[j]
		}
	}

	sort.Ints(k)
	ans := "No"
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idx := sort.Search(n*n, func(l int) bool { return kk[l] == m-k[i]-k[j] })
			if idx < len(k) && kk[idx] == m-k[i]-k[j] {
				ans = "Yes"
			}
		}
	}
	return ans
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
