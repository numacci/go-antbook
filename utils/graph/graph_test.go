package graph

import (
	"reflect"
	"testing"
)

func TestGraph_BellmanFord(t *testing.T) {
	type fields struct {
		V int
		E int
	}
	type args struct {
		es       []*Edge
		s        int
		directed bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
		want1  bool
	}{
		{name: "Bellman Ford Non Cyclic", fields: fields{
			V: 7,
			E: 10,
		}, args: args{
			es: []*Edge{
				{0, 2, 5},
				{0, 1, 2},
				{1, 2, 4},
				{1, 3, 6},
				{1, 4, 10},
				{2, 3, 2},
				{3, 5, 1},
				{4, 5, 3},
				{4, 6, 5},
				{5, 6, 9},
			},
			s: 0, directed: false,
		}, want: []int{0, 2, 5, 7, 11, 8, 16}, want1: false},
		{name: "Bellman Ford Cyclic", fields: fields{
			V: 6,
			E: 10,
		}, args: args{
			es: []*Edge{
				{0, 1, 5},
				{0, 2, 4},
				{1, 2, -2},
				{1, 3, 1},
				{2, 3, 2},
				{2, 4, 1},
				{2, 5, 4},
				{3, 1, -1},
				{3, 5, 3},
				{4, 5, 4},
			},
			s: 0, directed: true,
		}, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			_, got1 := g.BellmanFord(tt.args.es, tt.args.s, tt.args.directed)
			if got1 != tt.want1 {
				t.Errorf("BellmanFord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGraph_Dijkstra(t *testing.T) {
	type fields struct {
		V int
		E int
	}
	type args struct {
		G [][]*Edge
		s int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{name: "Dijkstra", fields: fields{
			V: 7,
			E: 10,
		}, args: args{
			G: [][]*Edge{
				{{To: 1, Cost: 2}, {To: 2, Cost: 5}},
				{{To: 0, Cost: 2}, {To: 2, Cost: 4}, {To: 3, Cost: 6}, {To: 4, Cost: 10}},
				{{To: 0, Cost: 5}, {To: 1, Cost: 4}, {To: 3, Cost: 2}},
				{{To: 1, Cost: 6}, {To: 2, Cost: 2}, {To: 5, Cost: 1}},
				{{To: 1, Cost: 10}, {To: 5, Cost: 3}, {To: 6, Cost: 5}},
				{{To: 3, Cost: 1}, {To: 4, Cost: 3}, {To: 6, Cost: 9}},
				{{To: 4, Cost: 5}, {To: 5, Cost: 9}},
			},
			s: 0,
		}, want: []int{0, 2, 5, 7, 11, 8, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if got := g.Dijkstra(tt.args.G, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_FloydWarshall(t *testing.T) {
	type fields struct {
		V int
		E int
	}
	type args struct {
		G [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "FloydWarshall", fields: fields{
			V: 7,
			E: 10,
		}, args: args{
			G: [][]int{
				{0, 2, 5, Inf, Inf, Inf, Inf},
				{2, 0, 4, 6, 10, Inf, Inf},
				{5, 4, 0, 2, Inf, Inf, Inf},
				{Inf, 6, 2, 0, Inf, 1, Inf},
				{Inf, 10, Inf, Inf, 0, 3, 5},
				{Inf, Inf, Inf, 1, 3, 0, 9},
				{Inf, Inf, Inf, Inf, 5, 9, 0},
			},
		}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				V: tt.fields.V,
				E: tt.fields.E,
			}
			if got := g.FloydWarshall(tt.args.G)[0][6]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FloydWarshall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_GetPath(t *testing.T) {
	type fields struct {
		V    int
		E    int
		Prev []int
	}
	type args struct {
		to int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{name: "GetPath", fields: fields{
			V:    7,
			E:    10,
			Prev: []int{-1, 0, 0, 2, 5, 3, 4},
		}, args: args{to: 6}, want: []int{0, 2, 3, 5, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				V:    tt.fields.V,
				E:    tt.fields.E,
				Prev: tt.fields.Prev,
			}
			if got := g.GetPath(tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
