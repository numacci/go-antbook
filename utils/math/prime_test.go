package math

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Yes", args: args{a: 53}, want: true},
		{name: "No", args: args{a: 295927}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.a); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivisor(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Prime", args: args{a: 53}, want: []int{1, 53}},
		{name: "NonPrime", args: args{a: 295927}, want: []int{1, 295927, 541, 547}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divisor(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorize(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{name: "PrimeFactor", args: args{a: 72}, want: map[int]int{2: 3, 3: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorize(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEratosthenes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{name: "Eratosthenes", args: args{n: 20}, want: 8, want1: []int{2, 3, 5, 7, 11, 13, 17, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Eratosthenes(tt.args.n)
			if got != tt.want {
				t.Errorf("Eratosthenes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Eratosthenes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSegmentEratosthenes(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{name: "SegmentEratosthenes", args: args{a: 22, b: 37}, want: 3, want1: []int{23, 29, 31}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SegmentEratosthenes(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SegmentEratosthenes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SegmentEratosthenes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
