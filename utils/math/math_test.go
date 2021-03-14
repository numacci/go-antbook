package math

import "testing"

func TestAbs(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Plus", args: args{3}, want: 3},
		{name: "Minus", args: args{-3}, want: 3},
		{name: "Zero", args: args{0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.a); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Max(t *testing.T) {
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
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Min(t *testing.T) {
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
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "GCD", args: args{a: 4, b: 6}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtGcd(t *testing.T) {
	type args struct {
		a int
		b int
		x *int
		y *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ExtGcd", args: args{
			a: 4,
			b: 11,
			x: new(int),
			y: new(int),
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtGcd(tt.args.a, tt.args.b, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ExtGcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
