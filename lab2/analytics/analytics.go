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

type Analytics struct {
	runInParallel   bool
	independentSets bool
}

func (a *Analytics) test5aWorker(fileName string, withRepetitions bool, wg *sync.WaitGroup) {
	var testType string

	if withRepetitions {
		testType = "with repetitions"
	} else {
		testType = "without repetitions"
	}

	fmt.Printf("Starting for %s\n", testType)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	k := 400

	algorithm := mc.New(sha256.New, k)

	for n := 1; n <= 10000; n++ {
		printProgress(testType, n, 1000)
		multiset := createDependentMultiset(n, 100000, withRepetitions)

		expected := countDistinct(multiset)
		estimated := algorithm.Count(multiset)

		_, err := w.WriteString(fmt.Sprintf("%t,%d,%d,%d\n", !withRepetitions, n, expected, estimated))
		checkError(err)

		w.Flush()
	}

	fmt.Printf("Done for %s\n", testType)
}

func (a *Analytics) Test5a(mainWg *sync.WaitGroup) {
	var wg sync.WaitGroup

	filesNames := []string{"data/5a_with_rep.csv", "data/5a_without_rep.csv"}

	if mainWg != nil {
		defer mainWg.Done()
	}

	wg.Add(2)

	go a.test5aWorker(filesNames[0], true, &wg)
	go a.test5aWorker(filesNames[1], false, &wg)

	wg.Wait()

	mergeFiles("data/5a.csv", "unique_multiset,n,expected,estimated\n", filesNames)
}

func (a *Analytics) test5bWorker(fileName string, k int, wg *sync.WaitGroup) {
	fmt.Printf("Starting for %d\n", k)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	algorithm := mc.New(sha256.New, k)

	var multiset []int

	last := 0

	for n := 1; n <= 10000; n++ {
		printProgress(strconv.Itoa(k), n, 1000)

		if a.independentSets {
			multiset, last = createIndependendMultiset(n, last+1)
		} else {
			multiset = createDependentMultiset(n, 100000, true)
		}

		var expected, estimated int

		expected = countDistinct(multiset)

		estimated = algorithm.Count(multiset)

		_, err := w.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", k, n, expected, estimated))
		checkError(err)

		w.Flush()
	}

	fmt.Printf("Done for %d\n", k)
}

func (a *Analytics) Test5b(mainWg *sync.WaitGroup) {
	var wg sync.WaitGroup
	var filesNames []string

	ks := []int{2, 3, 10, 100, 400}

	if mainWg != nil {
		defer mainWg.Done()
	}

	for _, k := range ks {
		wg.Add(1)

		fileName := fmt.Sprintf("data/5b_k_%d.csv", k)
		filesNames = append(filesNames, fileName)

		go a.test5bWorker(fileName, k, &wg)
	}

	wg.Wait()

	mergeFiles("data/5b.csv", "k,n,expected,estimated\n", filesNames)
}

func (a *Analytics) test5cWorker(fileName string, k int, wg *sync.WaitGroup) {
	fmt.Printf("Starting for %d\n", k)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	algorithm := mc.New(sha256.New, k)

	var multiset []int

	last := 0

	for n := 1; n <= 10000; n++ {
		printProgress(strconv.Itoa(k), n, 1000)

		if a.independentSets {
			multiset, last = createIndependendMultiset(n, last+1)
		} else {
			multiset = createDependentMultiset(n, 100000, true)
		}

		expected := countDistinct(multiset)
		estimated := algorithm.Count(multiset)

		_, err := w.WriteString(fmt.Sprintf("%d,%d,%d,%d\n", k, n, expected, estimated))
		checkError(err)

		w.Flush()
	}

	fmt.Printf("Done for %d\n", k)
}

func (a *Analytics) Test5c(mainWg *sync.WaitGroup) {
	var wg sync.WaitGroup
	var filesNames []string

	ks := []int{275, 280, 290}

	if mainWg != nil {
		defer mainWg.Done()
	}

	for _, k := range ks {
		wg.Add(1)

		fileName := fmt.Sprintf("data/5c_k_%d.csv", k)
		filesNames = append(filesNames, fileName)

		go a.test5cWorker(fileName, k, &wg)
	}

	wg.Wait()

	mergeFiles("data/5c.csv", "k,n,expected,estimated\n", filesNames)
}

func (a *Analytics) test6Worker(fileName, hashFuncName string, hashFunc func() hash.Hash, wg *sync.WaitGroup) {
	fmt.Printf("Starting for %s\n", hashFuncName)

	f, w := createFileWithWriter(fileName)
	defer f.Close()
	defer wg.Done()

	var multiset []int

	for b := 8; b <= 256; b <<= 1 {
		algorithm := mc.NewWithHashBitsLen(hashFunc, 400, b)

		printProgress(hashFuncName, b, 1)

		for i := 0; i < 1000; i++ {
			if a.independentSets {
				multiset, _ = createIndependendMultiset(10000, 0)
			} else {
				multiset = createDependentMultiset(10000, 100000, true)
			}

			expected := countDistinct(multiset)
			estimated := algorithm.Count(multiset)

			_, err := w.WriteString(fmt.Sprintf("%s,%d,%d,%d\n", hashFuncName, b, expected, estimated))
			checkError(err)
		}

		w.Flush()
	}

	fmt.Printf("Done for %s\n", hashFuncName)
}

func (a *Analytics) Test6(mainWg *sync.WaitGroup) {
	var wg sync.WaitGroup
	var filesNames []string

	hashes := map[string]func() hash.Hash{
		"md5":    md5.New,
		"sha1":   sha1.New,
		"sha256": sha256.New,
		"sha512": sha512.New,
	}

	if mainWg != nil {
		defer mainWg.Done()
	}

	for name, hashFunc := range hashes {
		wg.Add(1)

		fileName := fmt.Sprintf("data/6_hf_%s.csv", name)
		filesNames = append(filesNames, fileName)

		go a.test6Worker(fileName, name, hashFunc, &wg)
	}

	wg.Wait()

	mergeFiles("data/6.csv", "hash,b,expected,estimated\n", filesNames)
}

func (a *Analytics) TestAll() {
	var wg *sync.WaitGroup
	testFuncs := []func(*sync.WaitGroup){
		a.Test5a,
		a.Test5b,
		a.Test5c,
		a.Test6,
	}

	a.parseFlags()

	if !a.runInParallel {
		for i := range testFuncs {
			testFuncs[i](nil)
		}

		return
	}

	wg = &sync.WaitGroup{}

	wg.Add(len(testFuncs))

	for i := range testFuncs {
		go testFuncs[i](wg)
	}

	wg.Wait()
}
