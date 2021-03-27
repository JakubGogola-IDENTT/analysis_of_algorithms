package analytics

import (
	"crypto/sha256"
	"fmt"
	mc "lab2/mincount"
)

func Test5b() {
	ks := []int{2, 3, 10, 100, 400}

	f, w := createFileWithWriter("5a.csv")
	defer f.Close()

	_, err := w.WriteString("k,n,estimated_n\n")
	checkError(err)

	for _, k := range ks {
		algorithm := mc.New(sha256.New, k)

		fmt.Printf("k = %d\n", k)

		for n := 1; n <= 10000; n++ {
			printProgress(n, 100)
			multiset := createMultiset(n, 10000)

			expected := countDistinct(multiset)
			estimated := algorithm.Count(multiset)

			_, err := w.WriteString(fmt.Sprintf("%d,%d,%d\n", n, expected, estimated))
			checkError(err)
		}
		w.Flush()
	}
}
