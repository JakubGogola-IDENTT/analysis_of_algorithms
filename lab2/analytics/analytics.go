package analytics

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	mc "lab2/mincount"
	"strconv"
	"sync"
)

func test5bWorker(fileName string, k int, wg *sync.WaitGroup) {
	fmt.Printf("Starting for %d\n", k)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	algorithm := mc.New(sha256.New, k)

	for n := 1; n <= 10000; n++ {
		printProgress(strconv.Itoa(k), n, 1000)
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

func test6Worker(fileName, hashFuncName string, hashFunc func() hash.Hash, wg *sync.WaitGroup) {
	fmt.Printf("Starting for %s\n", hashFuncName)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	for b := 8; b <= 256; b <<= 1 {
		algorithm := mc.NewWithHashBitsLen(hashFunc, 400, b)

		printProgress(hashFuncName, b, 1)
		for i := 0; i < 1000; i++ {
			multiset := createMultiset(10000, 10000)

			expected := countDistinct(multiset)
			estimated := algorithm.Count(multiset)

			_, err := w.WriteString(fmt.Sprintf("%s,%d,%d,%d\n", hashFuncName, b, expected, estimated))
			checkError(err)
		}

		w.Flush()
	}

	fmt.Printf("Done for %s\n", hashFuncName)
}

func Test6() {
	var wg sync.WaitGroup
	var filesNames []string

	hashes := map[string]func() hash.Hash{
		"md5":    md5.New,
		"sha1":   sha1.New,
		"sha256": sha256.New,
		"sha512": sha512.New,
	}

	for name, hashFunc := range hashes {
		wg.Add(1)

		fileName := fmt.Sprintf("data/6_hf_%s.csv", name)
		filesNames = append(filesNames, fileName)

		go test6Worker(fileName, name, hashFunc, &wg)
	}

	wg.Wait()

	mergeFiles("data/6.csv", "hash,b,expected,estimated\n", filesNames)
}

func TestAll() {
	Test5b()
	Test6()
}
