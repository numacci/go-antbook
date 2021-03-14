package math

import "fmt"

func Abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func ExtGcd(a, b int, x, y *int) int {
	d := a
	if b != 0 {
		d = ExtGcd(b, a%b, y, x)
		*y -= (a / b) * *x
	} else {
		*x, *y = 1, 0
	}
	fmt.Printf("a=%v, b=%v, x=%v, y=%v\n", a, b, *x, *y)
	return d
}
