package ch3

import "testing"

func Test_p128(t *testing.T) {
	type args struct {
		n int
		k int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "lower_bound", args: args{n: 5, k: 3, a: []int{2, 3, 3, 5, 6}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p128(tt.args.n, tt.args.k, tt.args.a); got != tt.want {
				t.Errorf("p128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p129(t *testing.T) {
	type args struct {
		N int
		K int
		L []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Cable master", args: args{N: 4, K: 11, L: []float64{8.02, 7.43, 4.57, 5.39}}, want: 2.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p129(tt.args.N, tt.args.K, tt.args.L); got != tt.want {
				t.Errorf("p129() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p131(t *testing.T) {
	type args struct {
		N int
		M int
		x []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Aggressive cows", args: args{N: 5, M: 3, x: []int{1, 2, 8, 4, 9}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p131(tt.args.N, tt.args.M, tt.args.x); got != tt.want {
				t.Errorf("p131() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p132(t *testing.T) {
	type args struct {
		n int
		k int
		w []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Maximize average", args: args{
			n: 3,
			k: 2,
			w: []int{2, 5, 2},
			v: []int{2, 3, 1},
		}, want: 0.75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p132(tt.args.n, tt.args.k, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("p132() = %v, want %v", got, tt.want)
			}
		})
	}
}
