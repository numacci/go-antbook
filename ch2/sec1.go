package ch2

import (
	"github.com/numacci/go-antbook/utils/datastruct"
	"strings"
)

func p034(n int, k int, a []int) string {
	for bit := 0; bit < (1 << n); bit++ {
		sum := 0
		for i := 0; i < n; i++ {
			if (bit>>i)&1 == 1 {
				sum += a[i]
			}
		}
		if sum == k {
			return "Yes"
		}
	}
	return "No"
}

func dfs(i int, j int, N int, M int, G *[][]string) {
	(*G)[i][j] = "."

	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			ni, nj := i+di, j+dj
			if ni >= 0 && ni < N && nj >= 0 && nj < M && (*G)[ni][nj] == "W" {
				dfs(ni, nj, N, M, G)
			}
		}
	}
}

func p035(N int, M int, field []string) int {
	G := make([][]string, N)
	for i := 0; i < N; i++ {
		s := strings.Split(field[i], "")
		G[i] = s
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if G[i][j] == "W" {
				dfs(i, j, N, M, &G)
				ans++
			}
		}
	}
	return ans
}

const Inf = 1e8

func p037(N int, M int, field []string) int {
	// Covert []string to [][]string for simplicity
	G := make([][]string, N)
	for i := 0; i < N; i++ {
		G[i] = strings.Split(field[i], "")
	}

	var si, sj, gi, gj int
	d := make([][]int, N)
	// Get start/goal position and initialize distance slice
	for i := 0; i < N; i++ {
		d[i] = make([]int, M)
		for j := 0; j < M; j++ {
			if G[i][j] == "S" {
				si, sj = i, j
			}
			if G[i][j] == "G" {
				gi, gj = i, j
			}
			d[i][j] = Inf
		}
	}

	// Push start point to queue
	q := datastruct.NewPairQueue(N * M)
	q.Push(&datastruct.Pair{First: si, Second: sj})
	d[si][sj] = 0

	di, dj := []int{-1, 0, 1, 0}, []int{0, -1, 0, 1}
	for !q.IsEmpty() {
		p := q.Front()
		q.Pop()

		if p.First == gi && p.Second == gj {
			break
		}

		for i := 0; i < 4; i++ {
			ni := p.First + di[i]
			nj := p.Second + dj[i]

			inRange := ni >= 0 && ni < N && nj >= 0 && nj < M
			if inRange && G[ni][nj] != "#" && d[ni][nj] == Inf {
				q.Push(&datastruct.Pair{First: ni, Second: nj})
				d[ni][nj] = d[p.First][p.Second] + 1
			}
		}
	}
	return d[gi][gj]
}
