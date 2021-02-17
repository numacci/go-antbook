package datastruct

import (
	"fmt"
	"testing"
)

// Test_NotAllocate is performance test without allocating any capacity to slice.
func Test_NotAllocate(t *testing.T) {
	s := make([]int, 0)
	for i := 0; i < 1e9; i++ {
		s = append(s, i)
	}
	fmt.Println(len(s))
}

// Test_Allocate is performance test with allocating enough capacity to slice.
func Test_Allocate(t *testing.T) {
	s := make([]int, 0, 1e9)
	for i := 0; i < 1e9; i++ {
		s = append(s, i)
	}
	fmt.Println(len(s))
}

// Test_Initialize is performance test with initializing the elements of slice.
func Test_Initialize(t *testing.T) {
	s := make([]int, 1e9, 1e9)
	for i := 0; i < 1e9; i++ {
		s[i] = i
	}
	fmt.Println(len(s))
}
