package ch3

import (
	"reflect"
	"testing"
)

func Test_p135(t *testing.T) {
	type args struct {
		n int
		S int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Syakutori", args: args{
			n: 10,
			S: 15,
			a: []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
		}, want: 2},
		{name: "Syakutori", args: args{
			n: 5,
			S: 11,
			a: []int{1, 2, 3, 4, 5},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p135(tt.args.n, tt.args.S, tt.args.a); got != tt.want {
				t.Errorf("p135() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p137(t *testing.T) {
	type args struct {
		P int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Jessica's Reading Problem", args: args{P: 5, a: []int{1, 8, 8, 8, 1}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p137(tt.args.P, tt.args.a); got != tt.want {
				t.Errorf("p137() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p139(t *testing.T) {
	type args struct {
		N    int
		cows string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{name: "Face The Right Way", args: args{N: 7, cows: "BBFBFBB"}, want: 3, want1: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := p139(tt.args.N, tt.args.cows)
			if got != tt.want {
				t.Errorf("p139() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("p139() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_p145(t *testing.T) {
	type args struct {
		N int
		H int
		R int
		T int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "Physics Experiment", args: args{N: 1, H: 10, R: 10, T: 100}, want: []float64{4.95}},
		{name: "Physics Experiment", args: args{N: 2, H: 10, R: 10, T: 100}, want: []float64{4.95, 10.20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p145(tt.args.N, tt.args.H, tt.args.R, tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("p145() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p147(t *testing.T) {
	type args struct {
		n int
		A []int
		B []int
		C []int
		D []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "4 Values whose Sum is 0", args: args{
			n: 6,
			A: []int{-45, -41, -36, -36, 26, -32},
			B: []int{22, -27, 53, 30, -38, -54},
			C: []int{42, 56, -37, -75, -10, -6},
			D: []int{-16, 30, 77, -46, 62, 45},
		}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p147(tt.args.n, tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("p147() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p150(t *testing.T) {
	type args struct {
		w  int
		h  int
		n  int
		x1 []int
		x2 []int
		y1 []int
		y2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Compression", args: args{
			w:  10,
			h:  10,
			n:  5,
			x1: []int{1, 1, 4, 9, 10},
			x2: []int{6, 10, 4, 9, 10},
			y1: []int{4, 8, 1, 1, 6},
			y2: []int{4, 8, 10, 5, 10},
		}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p150(tt.args.w, tt.args.h, tt.args.n, tt.args.x1, tt.args.x2, tt.args.y1, tt.args.y2); got != tt.want {
				t.Errorf("p150() = %v, want %v", got, tt.want)
			}
		})
	}
}
