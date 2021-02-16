package ch1

import "testing"

func Test_p021(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// test cases
		{name: "Can create", args: args{n: 5, a: []int{2, 3, 4, 5, 10}}, want: 12},
		{name: "Cannot create", args: args{n: 4, a: []int{4, 5, 10, 20}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p021(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("p021() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p023(t *testing.T) {
	type args struct {
		L int
		n int
		x []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// test cases
		{
			name: "Ants",
			args: args{
				L: 10,
				n: 3,
				x: []int{2, 6, 7},
			},
			want:  4,
			want1: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := p023(tt.args.L, tt.args.n, tt.args.x)
			if got != tt.want {
				t.Errorf("p023() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("p023() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_p027(t *testing.T) {
	type args struct {
		n int
		m int
		k []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// test cases
		{name: "Possible", args: args{
			n: 3,
			m: 10,
			k: []int{1, 3, 5},
		}, want: "Yes"},
		{name: "Impossible", args: args{
			n: 3,
			m: 9,
			k: []int{1, 3, 5},
		}, want: "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p027(tt.args.n, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("p027() = %v, want %v", got, tt.want)
			}
		})
	}
}
