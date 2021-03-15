package ch2

import (
	"github.com/numacci/go-antbook/utils/math"
	math2 "math"
	"sort"
)

func p117(n int, v1, v2 []int) int {
	sort.Ints(v1)
	sort.Ints(v2)

	ans := 0
	for i := 0; i < n; i++ {
		ans += v1[i] * v2[n-i-1]
	}
	return ans
}

func p119(N int, mat [][]int) int {
	// 前処理：i行目における最後の1の位置をメモ
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = -1
		for j := 0; j < N; j++ {
			if mat[i][j] == 1 {
				a[i] = j
			}
		}
	}

	res := 0
	for i := 0; i < N; i++ {
		// i行目とSwapする行を探索
		pos := -1
		for j := i; j < N; j++ {
			if a[j] <= i {
				pos = j
				break
			}
		}

		for j := pos; j > i; j-- {
			a[j], a[j-1] = a[j-1], a[j]
			res++
		}
	}
	return res
}

func p121(P, Q int, A []int) int {
	A = append([]int{0}, A...)
	A = append(A, P+1)

	dp := make([][]int, Q+2) // (i, j) 番目の囚人を解放するために必要な最小金貨枚数
	for i := 0; i < Q+2; i++ {
		dp[i] = make([]int, Q+2)
	}

	// 幅 (解放する囚人数) が小さい順にdpを埋める
	for w := 2; w <= Q+1; w++ {
		for i := 0; i+w <= Q+1; i++ {
			j, t := i+w, int(1e9)

			for k := i + 1; k < j; k++ {
				t = math.Min(t, dp[i][k]+dp[k][j])
			}

			dp[i][j] = t + A[j] - A[i] - 2
		}
	}
	return dp[0][Q+1]
}

func p123(M, X int, P float64) float64 {
	n := 1 << M

	dp := make([][]float64, 2)
	prv, nxt := dp[0], dp[1]
	prv, nxt = make([]float64, n+1), make([]float64, n+1)

	prv[n] = 1.0
	for i := 0; i < M; i++ {
		for bit := 0; bit <= 1<<M; bit++ {
			jub, t := math.Min(bit, n-bit), 0.0
			for j := 0; j <= jub; j++ {
				t = math2.Max(t, P*prv[bit+j]+(1-P)*prv[bit-j])
			}
			nxt[bit] = t
		}
		prv, nxt = nxt, prv
	}
	return prv[X*n/1000000]
}
