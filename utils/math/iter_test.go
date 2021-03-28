package math

import (
	"reflect"
	"testing"
)

func TestLowerBound(t *testing.T) {
	type args struct {
		x   interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "[]int", args: args{
			x:   []int{1, 3, 4, 4, 6, 9},
			key: 4,
		}, want: 2},
		{name: "[]float64", args: args{
			x:   []float64{3.0, 4.0, 4.0, 6.0, 6.0, 9.0},
			key: 6.0,
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBound(tt.args.x, tt.args.key); got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	type args struct {
		x   interface{}
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "[]int", args: args{
			x:   []int{1, 3, 4, 4, 6, 9},
			key: 4,
		}, want: 4},
		{name: "[]float64", args: args{
			x:   []float64{3.0, 4.0, 4.0, 6.0, 6.0, 9.0},
			key: 6.0,
		}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperBound(tt.args.x, tt.args.key); got != tt.want {
				t.Errorf("UpperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInts(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Unique Int Slice", args: args{x: []int{1, 1, 3, 2, 3, 4}}, want: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInts(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
