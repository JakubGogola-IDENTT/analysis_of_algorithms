package analytics

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"lab3/hyperloglog"
	"sync"
)

type Analytics struct {
	minSize int
	maxSize int
}

func (a *Analytics) New(minSize, maxSize int) Analytics {
	return Analytics{
		minSize: minSize,
		maxSize: maxSize,
	}
}

func (a *Analytics) Init() {
	a.parseFlags()
}

func (a *Analytics) testWorker(fileName, hashFuncName string, hashFunc func() hash.Hash, wg *sync.WaitGroup) {
	fmt.Printf("Starting for %s hashing function\n", hashFuncName)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	var multiset []int
	next := 1

	for b := 4; b <= 16; b++ {
		hll := hyperloglog.New(hashFunc, b)

		for n := a.minSize; n <= a.maxSize; n++ {
			printProgress(hashFuncName, b, n, 1000)

			multiset, next = createMultiset(next, n)

			for _, v := range multiset {
				hll.Add(v)
			}

			expected := n
			estimated := hll.Count()

			_, err := w.WriteString(fmt.Sprintf("%s,%d,%d,%d,%d\n", hashFuncName, b, n, expected, estimated))
			checkError(err)

			hll.Clear()
		}

		w.Flush()
	}

	fmt.Printf("Done for %s hashing function\n", hashFuncName)
}

func (a *Analytics) Test() {
	var wg sync.WaitGroup
	var filesNames []string

	hashFuncs := map[string]func() hash.Hash{
		"md5":    md5.New,
		"sha1":   sha1.New,
		"sha256": sha256.New,
		"sha512": sha512.New,
	}

	for name, hashFunc := range hashFuncs {
		fileName := fmt.Sprintf("data/hll_%s.csv", name)
		filesNames = append(filesNames, fileName)

		wg.Add(1)
		go a.testWorker(fileName, name, hashFunc, &wg)
	}

	wg.Wait()

	mergeFiles("data/hll.csv", "hash,b,n,expected,estimated\n", filesNames)
}
