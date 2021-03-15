package ch2

import "testing"

func Test_p117(t *testing.T) {
	type args struct {
		n  int
		v1 []int
		v2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Minimum Scalar Product", args: args{
			n:  3,
			v1: []int{1, 3, -5},
			v2: []int{-2, 4, 1},
		}, want: -25},
		{name: "Minimum Scalar Product", args: args{
			n:  5,
			v1: []int{1, 2, 3, 4, 5},
			v2: []int{1, 0, 1, 0, 1},
		}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p117(tt.args.n, tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("p117() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p119(t *testing.T) {
	type args struct {
		N   int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Crazy Rows", args: args{
			N: 2,
			mat: [][]int{
				{1, 0},
				{1, 1},
			},
		}, want: 0},
		{name: "Crazy Rows", args: args{
			N: 3,
			mat: [][]int{
				{0, 0, 1},
				{1, 0, 0},
				{0, 1, 0},
			},
		}, want: 2},
		{name: "Crazy Rows", args: args{
			N: 4,
			mat: [][]int{
				{1, 1, 1, 0},
				{1, 1, 0, 0},
				{1, 1, 0, 0},
				{1, 0, 0, 0},
			},
		}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p119(tt.args.N, tt.args.mat); got != tt.want {
				t.Errorf("p119() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p121(t *testing.T) {
	type args struct {
		P int
		Q int
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Bribe the Prisoners", args: args{P: 8, Q: 1, A: []int{3}}, want: 7},
		{name: "Bribe the Prisoners", args: args{P: 20, Q: 3, A: []int{3, 6, 14}}, want: 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p121(tt.args.P, tt.args.Q, tt.args.A); got != tt.want {
				t.Errorf("p121() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p123(t *testing.T) {
	type args struct {
		M int
		X int
		P float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Millionaire", args: args{M: 1, X: 500000, P: 0.5}, want: 0.5},
		{name: "Millionaire", args: args{M: 3, X: 600000, P: 0.75}, want: 0.843750},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p123(tt.args.M, tt.args.X, tt.args.P); got != tt.want {
				t.Errorf("p123() = %v, want %v", got, tt.want)
			}
		})
	}
}
