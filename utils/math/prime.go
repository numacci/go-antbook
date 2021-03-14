package math

import (
	"math"
)

func IsPrime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func Divisor(a int) []int {
	res := make([]int, 0, a)
	for i := 1; i*i <= a; i++ {
		if a%i == 0 {
			res = append(res, i)
			if i != a/i {
				res = append(res, a/i)
			}
		}
	}
	return res
}

func Factorize(a int) map[int]int {
	res := make(map[int]int)
	for i := 2; i*i <= a; i++ {
		for a%i == 0 {
			res[i]++
			a /= i
		}
	}
	if a != 1 {
		res[a] = 1
	}
	return res
}

func Eratosthenes(n int) (int, []int) {
	p, prime := 0, make([]int, 0, n)

	isPrime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		isPrime[i] = true
	}

	isPrime[0], isPrime[1] = false, false
	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		prime = append(prime, i)
		p++
		// 発見された素数の倍数を除外する
		for j := 2 * i; j <= n; j += i {
			isPrime[j] = false
		}
	}
	return p, prime
}

func SegmentEratosthenes(a, b int) (int, []int) {
	isPrimeAtoB := make([]bool, b-a)                             // [a, b) の櫛
	isPrimeToSqrtB := make([]bool, int(math.Sqrt(float64(b)))+1) // [2, sqrt(b)) の櫛

	for i := 0; i*i < b; i++ {
		isPrimeToSqrtB[i] = true
	}
	for i := 0; i < b-a; i++ {
		isPrimeAtoB[i] = true
	}

	for i := 2; i*i < b; i++ {
		if !isPrimeToSqrtB[i] {
			continue
		}
		for j := 2 * i; j*j < b; j += i {
			isPrimeToSqrtB[j] = false // [2, sqrt(b)) の櫛
		}
		for j := Max(i, (a+i-1)/i) * i; j < b; j += i {
			isPrimeAtoB[j-a] = false // [a, b) の櫛
		}
	}

	p, prime := 0, make([]int, 0, b-a)
	for i := a; i < b; i++ {
		if isPrimeAtoB[i-a] {
			p++
			prime = append(prime, i)
		}
	}
	return p, prime
}
