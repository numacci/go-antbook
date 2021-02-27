package ch2

import "testing"

func Test_p052(t *testing.T) {
	type args struct {
		n int
		W int
		w []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "01 knapsack", args: args{
			n: 4,
			W: 5,
			w: []int{2, 1, 3, 2},
			v: []int{3, 2, 4, 2},
		}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p052(tt.args.n, tt.args.W, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("p052() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p056(t *testing.T) {
	type args struct {
		n int
		m int
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "LCS", args: args{
			n: 4,
			m: 4,
			s: "abcd",
			t: "becd",
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p056(tt.args.n, tt.args.m, tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("p056() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p058(t *testing.T) {
	type args struct {
		n int
		W int
		w []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "unlimited knapsack", args: args{
			n: 3,
			W: 7,
			w: []int{3, 4, 2},
			v: []int{4, 5, 3},
		}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p058(tt.args.n, tt.args.W, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("p058() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p060(t *testing.T) {
	type args struct {
		n int
		W int
		w []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "01 knapsack", args: args{
			n: 4,
			W: 5,
			w: []int{2, 1, 3, 2},
			v: []int{3, 2, 4, 2},
		}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p060(tt.args.n, tt.args.W, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("p060() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p062(t *testing.T) {
	type args struct {
		n int
		K int
		a []int
		m []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Limited Sum of Subsequence", args: args{
			n: 3,
			K: 17,
			a: []int{3, 5, 8},
			m: []int{3, 2, 2},
		}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p062(tt.args.n, tt.args.K, tt.args.a, tt.args.m); got != tt.want {
				t.Errorf("p062() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p063(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "LIS", args: args{
			n: 5,
			a: []int{4, 2, 3, 1, 5},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p063(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("p063() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p066(t *testing.T) {
	type args struct {
		n int
		m int
		M int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Div Num", args: args{n: 4, m: 3, M: 10000}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p066(tt.args.n, tt.args.m, tt.args.M); got != tt.want {
				t.Errorf("p066() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p067(t *testing.T) {
	type args struct {
		n int
		m int
		M int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Combination with repetition", args: args{
			n: 3,
			m: 3,
			M: 10000,
			a: []int{1, 2, 3},
		}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p067(tt.args.n, tt.args.m, tt.args.M, tt.args.a); got != tt.want {
				t.Errorf("p067() = %v, want %v", got, tt.want)
			}
		})
	}
}
