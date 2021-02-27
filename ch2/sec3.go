package ch2

import (
	"github.com/numacci/go-antbook/utils/math"
	"sort"
)

func p052(n, W int, w, v []int) int {
	// dp[i][j]: i番目までの荷物を利用する場合に，重さの総和がjを超えないような価値の最大値
	dp := make([][]int, n+1)
	// memset
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+1)
		if i == 0 {
			continue
		}
		for j := 1; j <= W; j++ {
			dp[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= W; j++ {
			if j >= w[i] {
				dp[i+1][j] = math.Max(dp[i][j], dp[i][j-w[i]]+v[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	return dp[n][W]
}

func p056(n, m int, s, t string) int {
	// dp[i][j]: s1...siとt1...tjに対するLCSの長さ
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	rs, rt := []rune(s), []rune(t)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rs[i] == rt[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = math.Max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[n][m]
}

func p058(n, W int, w, v []int) int {
	dp := make([]int, W+1)
	for i := 0; i < n; i++ {
		for j := w[i]; j <= W; j++ {
			dp[j] = math.Max(dp[j], dp[j-w[i]]+v[i])
		}
	}
	return dp[W]
}

const (
	MAXV = 100*100 + 1
	INF  = 1e9
)

func p060(n, W int, w, v []int) int {
	// dp[i][j]: i番目までの品物を利用する場合に，価値がjとなる最小の重さ
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, MAXV)
		for j := 0; j < MAXV; j++ {
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j < MAXV; j++ {
			if j < v[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = math.Min(dp[i][j], dp[i][j-v[i]]+w[i])
			}
		}
	}

	ans := 0
	for i := 1; i < MAXV; i++ {
		if dp[n][i] <= W {
			ans = i
		}
	}
	return ans
}

func p062(n, K int, a, m []int) string {
	// dp[i][j]: i番目まででjを作る際に余るi番目のaの残数
	dp := make([]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= K; j++ {
			if dp[j] >= 0 {
				dp[j] = m[i]
			} else if j < a[i] || dp[j-a[i]] <= 0 {
				dp[j] = -1
			} else {
				dp[j] = dp[j-a[i]] - 1
			}
		}
	}
	if dp[K] >= 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func p063(n int, a []int) int {
	// dp[i]:長さがiであるような部分増加列における最後の要素
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = Inf
	}

	for i := 0; i < n; i++ {
		idx := sort.Search(n+1, func(j int) bool { return a[i] < dp[j] })
		dp[idx] = a[i]
	}
	ans := 0
	for i := 0; i <= n; i++ {
		if dp[i] == Inf {
			break
		}
		ans++
	}
	return ans
}

const (
	Mod = 1e9 + 7
)

func p066(n, m, M int) int {
	// dp[i][j]: jのi分割の総数
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j-i >= 0 {
				dp[i][j] = (dp[i-1][j] + dp[i][j-i]) % Mod
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}

func p067(n, m, M int, a []int) int {
	// dp[i][j]: i番目の品物からj個選ぶ組み合わせの総数
	dp := make([][]int, n+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if j-1-a[i] >= 0 {
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j] - dp[i][j-1-a[i]] + M) % M
			} else {
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j]) % M
			}
		}
	}
	return dp[n][m]
}
