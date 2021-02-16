package utils

import "testing"

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Plus sign", args: args{a: 1, b: 3}, want: 3},
		{name: "Minus sign", args: args{a: -1, b: -3}, want: -1},
		{name: "Different sign", args: args{a: -1, b: 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Plus sign", args: args{a: 1, b: 3}, want: 1},
		{name: "Minus sign", args: args{a: -1, b: -3}, want: -3},
		{name: "Different sign", args: args{a: -1, b: 3}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
