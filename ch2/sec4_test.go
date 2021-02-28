package ch2

import "testing"

func Test_p073(t *testing.T) {
	type args struct {
		N int
		L int
		P int
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Expedition", args: args{
			N: 4,
			L: 25,
			P: 10,
			A: []int{10, 14, 20, 21},
			B: []int{10, 5, 2, 4},
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p073(tt.args.N, tt.args.L, tt.args.P, tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("p073() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p085(t *testing.T) {
	type args struct {
		N int
		K int
		T []int
		X []int
		Y []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Food chain", args: args{
			N: 100,
			K: 7,
			T: []int{1, 2, 2, 2, 1, 2, 1},
			X: []int{101, 1, 2, 3, 1, 3, 5},
			Y: []int{1, 2, 3, 3, 3, 1, 5},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p085(tt.args.N, tt.args.K, tt.args.T, tt.args.X, tt.args.Y); got != tt.want {
				t.Errorf("p085() = %v, want %v", got, tt.want)
			}
		})
	}
}
