package ch2

import "testing"

func Test_p042(t *testing.T) {
	type args struct {
		c []int
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// test cases
		{name: "Coins", args: args{
			c: []int{3, 2, 1, 3, 0, 2},
			A: 620,
		}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p042(tt.args.c, tt.args.A); got != tt.want {
				t.Errorf("p042() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p043(t *testing.T) {
	type args struct {
		n int
		s []int
		t []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// test cases
		{name: "Scheduling", args: args{
			n: 5,
			s: []int{1, 2, 4, 6, 8},
			t: []int{3, 5, 7, 9, 10},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p043(tt.args.n, tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("p043() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p045(t *testing.T) {
	type args struct {
		N int
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Best Cow Line", args: args{N: 6, S: "ACDBCB"}, want: "ABCBCD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p045(tt.args.N, tt.args.S); got != tt.want {
				t.Errorf("p045() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p047(t *testing.T) {
	type args struct {
		N int
		R int
		X []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Saruman's Army", args: args{
			N: 6,
			R: 10,
			X: []int{1, 7, 15, 20, 30, 50},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p047(tt.args.N, tt.args.R, tt.args.X); got != tt.want {
				t.Errorf("p047() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p049(t *testing.T) {
	type args struct {
		N int
		L []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Fence Repair1", args: args{N: 3, L: []int{8, 5, 8}}, want: 34},
		{name: "Fence Repair2", args: args{N: 5, L: []int{3, 4, 5, 1, 2}}, want: 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p049(tt.args.N, tt.args.L); got != tt.want {
				t.Errorf("p049() = %v, want %v", got, tt.want)
			}
		})
	}
}
