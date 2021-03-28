package datastruct

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	// int
	intSet := NewSet(10)
	fmt.Println(intSet)

	intSet.Insert(10)
	if !intSet.Contains(10) {
		t.Fatalf("True is expected, but %v\n", intSet.Contains(10))
	}
	fmt.Println(intSet)
	intSet.Insert(-3)
	if !intSet.Contains(-3) {
		t.Fatalf("True is expected, but %v\n", intSet.Contains(10))
	}
	fmt.Println(intSet)

	// duplicated
	intSet.Insert(-3)
	if intSet.Size() != 2 {
		t.Fatalf("2 is expected, but %v\n", intSet.Size())
	}
	fmt.Println(intSet)

	intSet.Erase(-3)
	if intSet.Contains(-3) {
		t.Fatalf("True is expected, but %v\n", intSet.Contains(-3))
	}
	fmt.Println(intSet)
	intSet.Erase(10)
	if intSet.Contains(10) {
		t.Fatalf("True is expected, but %v\n", intSet.Contains(10))
	}
	fmt.Println(intSet)

	// string
	strSet := NewSet(10)
	fmt.Println(strSet)

	strSet.Insert("abc")
	if !strSet.Contains("abc") {
		t.Fatalf("True is expected, but %v\n", strSet.Contains("abc"))
	}
	fmt.Println(strSet)
	strSet.Insert("def")
	if !strSet.Contains("def") {
		t.Fatalf("True is expected, but %v\n", strSet.Contains("def"))
	}
	fmt.Println(strSet)

	// duplicated
	strSet.Insert("def")
	if strSet.Size() != 2 {
		t.Fatalf("2 is expected, but %v\n", strSet.Size())
	}
	fmt.Println(strSet)

	strSet.Erase("abc")
	if strSet.Contains("abc") {
		t.Fatalf("True is expected, but %v\n", strSet.Contains("abc"))
	}
	fmt.Println(strSet)
}
