package math

func NextPermutation(n []int) bool {
	i := len(n) - 2
	if i < 0 {
		return false
	}

	for ; n[i] >= n[i+1]; i-- {
		// completed
		if i == 0 {
			return false
		}
	}

	j := len(n) - 1
	for n[i] >= n[j] {
		j--
	}
	// swap
	n[i], n[j] = n[j], n[i]

	// reverse
	l, r := i+1, len(n)-1
	for l < r {
		n[l], n[r] = n[r], n[l]
		l++
		r--
	}
	return true
}
