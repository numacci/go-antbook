package ch2

import "testing"

func Test_p034(t *testing.T) {
	type args struct {
		n int
		k int
		a []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// test cases
		{name: "Possible", args: args{
			n: 4,
			k: 13,
			a: []int{1, 2, 4, 7},
		}, want: "Yes"},
		{name: "Impossible", args: args{
			n: 4,
			k: 15,
			a: []int{1, 2, 4, 7},
		}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p034(tt.args.n, tt.args.k, tt.args.a); got != tt.want {
				t.Errorf("p034() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p035(t *testing.T) {
	type args struct {
		N     int
		M     int
		field []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// test cases
		{name: "Lake counting", args: args{
			N: 10,
			M: 12,
			field: []string{
				"W........WW.",
				".WWW.....WWW",
				"....WW...WW.",
				".........WW.",
				".........W..",
				"..W......W..",
				".W.W.....WW.",
				"W.W.W.....W.",
				".W.W......W.",
				"..W.......W.",
			},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p035(tt.args.N, tt.args.M, tt.args.field); got != tt.want {
				t.Errorf("p035() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p037(t *testing.T) {
	type args struct {
		N     int
		M     int
		field []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// test cases
		{name: "BFS", args: args{
			N: 10,
			M: 10,
			field: []string{
				"#S######.#",
				"......#..#",
				".#.##.##.#",
				".#........",
				"##.##.####",
				"....#....#",
				".#######.#",
				"....#.....",
				".####.###.",
				"....#...G#",
			},
		}, want: 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p037(tt.args.N, tt.args.M, tt.args.field); got != tt.want {
				t.Errorf("p037() = %v, want %v", got, tt.want)
			}
		})
	}
}
