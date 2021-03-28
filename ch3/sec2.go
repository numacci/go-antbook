package ch3

import (
	"github.com/numacci/go-antbook/utils/datastruct"
	"github.com/numacci/go-antbook/utils/math"
	math2 "math"
	"sort"
)

func p135(n, S int, a []int) int {
	// Solution1. binary search
	//sum := make([]int, n+1)
	//for i := 0; i < n; i++ {
	//	sum[i+1] = sum[i] + a[i]
	//}
	//
	//if sum[n] < S {
	//	return 0
	//}
	//
	//// [s, t) の区間和がS以上か二分探索で求める
	//res := n
	//for s := 1; sum[n]-sum[s] >= S; s++ {
	//	t := math.LowerBound(sum[s:], sum[s]+S)
	//	res = math.Min(res, t)
	//}
	//return res

	// Solution2. Shakutori
	s, t, sum := 0, 0, 0
	res := n + 1
	for {
		for t < n && sum < S {
			sum += a[t]
			t++
		}

		if sum < S {
			break
		}

		res = math.Min(res, t-s)
		sum -= a[s]
		s++
	}
	return res
}

func p137(P int, a []int) int {
	set := datastruct.NewSet(len(a))
	for _, x := range a {
		set.Insert(x)
	}
	n := set.Size()

	s, t, num := 0, 0, 0
	count := make(map[int]int, set.Size()) // ページ番号：出現数
	res := P
	for {
		for t < P && num < n {
			v, b := count[a[t]]
			if v == 0 || !b {
				num++
			}
			count[a[t]]++
			t++
		}

		if num < n {
			break
		}

		res = math.Min(res, t-s)
		count[a[s]]--
		if count[a[s]] == 0 {
			num--
		}
		s++
	}
	return res
}

func p139(N int, cows string) (int, int) {
	dir := make([]int, 0, len(cows))
	for _, b := range cows {
		if b == 'F' {
			dir = append(dir, 0)
		} else {
			dir = append(dir, 1)
		}
	}

	K, M := 1, N
	for k := 1; k <= N; k++ {
		m := calc(N, k, dir)
		if m >= 0 && M > m {
			M = m
			K = k
		}
	}
	return K, M
}

// Kを固定したときの最小操作回数mを求める
func calc(N, K int, dir []int) int {
	f := make([]int, 5*1e3) // 区間 [i, i+K-1] を反転させたかどうか
	res, sum := 0, 0        // sum: fの和
	for i := 0; i+K <= N; i++ {
		if (dir[i]+sum)%2 != 0 {
			// 先頭の牛が後ろを向いている
			res++
			f[i] = 1
		}
		sum += f[i]
		if i-K+1 >= 0 {
			sum -= f[i-K+1] // 加算済みで不要な値 (範囲外) をリセット
		}
	}

	// 余りの牛が前を向いているかチェック
	for i := N - K + 1; i < N; i++ {
		if (dir[i]+sum)%2 != 0 {
			return -1 // 解なし
		}
		if i-K+1 >= 0 {
			sum -= f[i-K+1]
		}
	}
	return res
}

func p145(N, H, R, T int) []float64 {
	pos := make([]float64, N)
	g := 10.0

	for i := 0; i < N; i++ {
		t := float64(T - i)
		if t < 0 {
			pos[i] = float64(H)
			continue
		}

		gt := math2.Sqrt(2 * float64(H) / g)
		k, dt := int(t/gt), 0.0
		if k%2 == 0 {
			// 落下中
			dt = t - float64(k)*gt
		} else {
			// 上昇中
			dt = gt - (t - float64(k)*gt)
		}
		pos[i] = float64(H) - g*dt*dt/float64(2)
	}

	sort.Float64s(pos)

	for i := 0; i < N; i++ {
		pos[i] = pos[i] + 2*float64(R)*float64(i)/100.0
		pos[i] = math2.Round(pos[i]*100.0) / 100.0
	}
	return pos
}

func p147(n int, A, B, C, D []int) int {
	CD := make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			CD[n*i+j] = C[i] + D[j]
		}
	}
	sort.Ints(CD)

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cd := -(A[i] + B[j])
			res += math.UpperBound(CD, cd) - math.LowerBound(CD, cd)
		}
	}
	return res
}

func compress(x1, x2 []int, w int) int {
	n := len(x1)
	xs := make([]int, 0, 4*n)

	for i := 0; i < n; i++ {
		for d := -1; d <= 0; d++ {
			// 棒の端点-1~+1の座標を残す
			tx1, tx2 := x1[i]+d, x2[i]+d
			if 1 <= tx1 && tx1 <= w {
				xs = append(xs, tx1)
			}
			if 1 <= tx2 && tx2 <= w {
				xs = append(xs, tx2)
			}
		}
	}

	// 重複を削除→圧縮後に存在すべきオリジナルの座標が残る
	xs = math.UniqueInts(xs)
	for i := 0; i < n; i++ {
		// 圧縮後残る座標内でオリジナルの座標が何番目に位置するか→圧縮された座標が得られる
		x1[i] = sort.Search(len(xs), func(j int) bool { return x1[i] <= xs[j] })
		x2[i] = sort.Search(len(xs), func(j int) bool { return x2[i] <= xs[j] })
	}
	return len(xs)
}

func p150(w, h, n int, x1, x2, y1, y2 []int) int {
	// 座標圧縮
	w = compress(x1, x2, w)
	h = compress(y1, y2, h)

	// 塗りつぶし
	field := make([][]bool, h)
	for i := 0; i < h; i++ {
		field[i] = make([]bool, w)
	}
	for i := 0; i < n; i++ {
		for y := y1[i]; y <= y2[i]; y++ {
			for x := x1[i]; x <= x2[i]; x++ {
				field[y][x] = true
			}
		}
	}

	// 領域数え上げ
	ans := 0
	dx, dy := []int{-1, 0, 1, 0}, []int{0, -1, 0, 1}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if field[y][x] {
				continue
			}

			field[y][x] = true
			ans++

			// 幅優先探索で塗り潰す
			que := datastruct.NewPairQueue(w * h)
			que.Push(&datastruct.Pair{First: x, Second: y})
			for !que.IsEmpty() {
				sx, sy := que.Front().First, que.Front().Second
				que.Pop()

				for i := 0; i < 4; i++ {
					tx, ty := sx+dx[i], sy+dy[i]
					if tx < 0 || tx >= w || ty < 0 || ty >= h || field[ty][tx] {
						continue
					}
					que.Push(&datastruct.Pair{First: tx, Second: ty})
					field[ty][tx] = true
				}

			}
		}
	}
	return ans
}
