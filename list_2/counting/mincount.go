package counting

import "sort"

// counting is implementation of Min Count algorithm
// multiset represents set of values to count.
// k is size of hashes array
func counting(multiset []int, k int) float64 {
	var M = ones(k)

	k--

	for _, x := range multiset {
		h := getHash(x)

		if h < M[k] && !includes(h, M) {
			M[k] = h
			sort.Float64s(M)
		}
	}

	if M[k] == 1. {
		for _, m := range M {
			if m != 1. {
				return m
			}
		}
	}

	return float64(k-1) / M[k]
}
