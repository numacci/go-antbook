package math

import "testing"

func TestModPow(t *testing.T) {
	type args struct {
		x   int
		n   int
		mod int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ModPow", args: args{x: 2, n: 4, mod: 4}, want: 0},
		{name: "ModPow", args: args{x: 3, n: 4, mod: 4}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModPow(tt.args.x, tt.args.n, tt.args.mod); got != tt.want {
				t.Errorf("ModPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
