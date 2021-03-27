package analytics

import (
	"crypto/sha256"
	"fmt"
	mc "lab2/mincount"
	"sync"
)

func test5bWorker(fileName string, k int, wg *sync.WaitGroup) {
	fmt.Printf("Starting for %d\n", k)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	algorithm := mc.New(sha256.New, k)

	for n := 1; n <= 10000; n++ {
		multiset := createMultiset(n, 10000)

		expected := countDistinct(multiset)
		estimated := algorithm.Count(multiset)

		_, err := w.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", k, n, expected, estimated))
		checkError(err)

		w.Flush()
	}

	fmt.Printf("Done for %d\n", k)
}

func Test5b() {
	var wg sync.WaitGroup
	var filesNames []string

	ks := []int{2, 3, 10, 100, 400}

	for _, k := range ks {
		wg.Add(1)

		fileName := fmt.Sprintf("data/5b_k_%d.csv", k)
		filesNames = append(filesNames, fileName)

		go test5bWorker(fileName, k, &wg)
	}

	wg.Wait()

	mergeFiles("data/5b.csv", "k,n,expected,estimated\n", filesNames)
}

func TestAll() {
	Test5b()
}
