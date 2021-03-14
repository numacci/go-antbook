package ch2

import (
	"github.com/numacci/go-antbook/utils/math"
)

func p107(P1, P2 []int) int {
	x := math.Abs(P1[0] - P2[0])
	y := math.Abs(P1[1] - P2[1])

	if x == 0 && y == 0 {
		return 0
	}

	gcd := math.Gcd(x, y)
	return gcd - 1
}

// Refer to go-antbook/utils/math directory to see other implementation of this section.
