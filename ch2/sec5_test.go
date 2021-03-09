package ch2

import (
	"github.com/numacci/go-antbook/utils/graph"
	"testing"
)

func Test_p093(t *testing.T) {
	type args struct {
		n   int
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Binary Graph (No)", args: args{
			n:   3,
			mat: [][]int{{0, 1}, {1, 2}, {2, 0}},
		}, want: "No"},
		{name: "Binary Graph (Yes)", args: args{
			n:   4,
			mat: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
		}, want: "Yes"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p093(tt.args.n, tt.args.mat); got != tt.want {
				t.Errorf("p093() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p101(t *testing.T) {
	type args struct {
		V  int
		E  int
		es []*graph.Edge
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Kruskal", args: args{
			V: 7,
			E: 9,
			es: []*graph.Edge{
				{0, 1, 10},
				{0, 3, 2},
				{1, 4, 5},
				{2, 3, 1},
				{3, 4, 7},
				{3, 5, 3},
				{4, 5, 1},
				{4, 6, 8},
				{5, 6, 5},
			},
		}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p101(tt.args.V, tt.args.E, tt.args.es); got != tt.want {
				t.Errorf("p101() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p102(t *testing.T) {
	type args struct {
		N  int
		R  int
		gh [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Roadblocks", args: args{
			N: 4,
			R: 4,
			gh: [][]int{
				{1, 2, 100},
				{2, 3, 250},
				{2, 4, 200},
				{3, 4, 100},
			},
		}, want: 450},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p102(tt.args.N, tt.args.R, tt.args.gh); got != tt.want {
				t.Errorf("p102() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p103(t *testing.T) {
	type args struct {
		N   int
		M   int
		R   int
		rel [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Conscription", args: args{
			N: 5,
			M: 5,
			R: 8,
			rel: [][]int{
				{4, 3, 6831},
				{1, 3, 4583},
				{0, 0, 6592},
				{0, 1, 3063},
				{3, 3, 4975},
				{1, 3, 2049},
				{4, 2, 2104},
				{2, 2, 781},
			},
		}, want: 71071},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p103(tt.args.N, tt.args.M, tt.args.R, tt.args.rel); got != tt.want {
				t.Errorf("p103() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p104(t *testing.T) {
	type args struct {
		N  int
		ML int
		MD int
		ok [][]int
		ng [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Layout", args: args{
			N:  4,
			ML: 2,
			MD: 1,
			ok: [][]int{
				{1, 3, 10},
				{2, 4, 20},
			},
			ng: [][]int{
				{2, 3, 3},
			},
		}, want: 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p104(tt.args.N, tt.args.ML, tt.args.MD, tt.args.ok, tt.args.ng); got != tt.want {
				t.Errorf("p104() = %v, want %v", got, tt.want)
			}
		})
	}
}
