package datastruct

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(7)
	fmt.Println(uf)

	uf.Unite(3, 5)
	if !uf.Same(3, 5) {
		t.Errorf("True is expected, but False for uf.Same(3, 5)\n")
	}
	fmt.Println(uf)
	uf.Unite(2, 3)
	if !uf.Same(2, 3) {
		t.Errorf("True is expected, but False for uf.Same(2, 3)\n")
	}
	fmt.Println(uf)
	uf.Unite(4, 5)
	if !uf.Same(4, 5) {
		t.Errorf("True is expected, but False for uf.Same(4, 5)\n")
	}
	fmt.Println(uf)

	if uf.Same(1, 5) {
		t.Errorf("False is expected, but True for uf.Same(1, 5)\n")
	}
}
