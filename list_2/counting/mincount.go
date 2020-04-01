package counting

import (
	"fmt"
	"log"
	"os"
	"sort"
)

// counting is implementation of Min Count algorithm
// multiset represents set of values to count.
// k is size of hashes array
func counting(multiset []int, k int) int {
	var M = ones(k)

	for _, x := range multiset {
		h := getHash(x)

		if h < M[k-1] && !includes(h, M) {
			M[k-1] = h
			sort.Float64s(M)
		}
	}

	if M[k-1] == 1. {
		return countNonOnes(M)
	}

	return int(float64(k-1) / M[k-1])
}

func runTests(withRepetitions bool) {
	ks := []int{2, 3, 10, 100, 400}

	for _, k := range ks {
		test(k, withRepetitions)
	}

}

func test(k int, withRepetitions bool) {
	file, err := os.Create(getFileName(k, withRepetitions))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Fprintf(file, "k,n,n_wave\n")

	fmt.Printf("Testing for k = %d\n", k)

	for n := 1; n <= 10000; n++ {
		var multiset []int

		if withRepetitions {
			multiset = generateMultisetWithRepetitions(n)
		} else {
			multiset = generateMultiset(n)
		}

		estimation := counting(multiset, k)

		fmt.Fprintf(file, "%d,%d,%d\n", k, n, estimation)
	}

	fmt.Printf("End test for k = %d\n", k)
}
