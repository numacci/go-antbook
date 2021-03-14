package ch2

import "testing"

func Test_p107(t *testing.T) {
	type args struct {
		P1 []int
		P2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "LatticePoint", args: args{
			P1: []int{1, 11},
			P2: []int{5, 3},
		}, want: 3},
		{name: "Horizontal", args: args{
			P1: []int{1, 11},
			P2: []int{1, 3},
		}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p107(tt.args.P1, tt.args.P2); got != tt.want {
				t.Errorf("p107() = %v, want %v", got, tt.want)
			}
		})
	}
}
