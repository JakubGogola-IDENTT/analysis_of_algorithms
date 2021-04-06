package analytics

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
)

func (a *Analytics) parseFlags() {
	flag.BoolVar(&a.runInParallel, "p", false, "Run all tests in parallel")
	flag.BoolVar(&a.independentSets, "i", true, "Use indepndent sets for tests")
	flag.Parse()
}

func createFileWithWriter(fileName string) (*os.File, *bufio.Writer) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)

	return f, w
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printProgress(param string, v, div int) {
	if v%div == 0 {
		fmt.Printf("%s | Progress: %d...\n", param, v)
	}
}

func getRandomInt(size int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(size)))

	if err != nil {
		log.Fatal(err)
	}

	return int(n.Int64())
}

func countDistinct(vs []int) int {
	distinctVs := make(map[int]bool)

	for _, v := range vs {
		distinctVs[v] = true
	}

	return len(distinctVs)
}

func createDependentMultiset(size, randRange int, withRepetitions bool) (multiset []int) {
	usedValues := make(map[int]bool)
	multiset = make([]int, size)

	i := 0
	for i < size {
		n := getRandomInt(randRange)

		if !withRepetitions && usedValues[n] {
			continue
		}

		usedValues[n] = true
		multiset[i] = n
		i++
	}

	return multiset
}

func createIndependendMultiset(size int) (multiset []int) {
	multiset = make([]int, size)

	for i := 0; i < size; i++ {
		multiset[i] = i + 1
	}

	return multiset
}

func mergeFiles(mainFileName, columns string, filesNames []string) {
	mf, w := createFileWithWriter(mainFileName)
	defer mf.Close()

	_, err := w.WriteString(columns)
	checkError(err)

	for _, name := range filesNames {
		f, err := os.Open(name)
		checkError(err)

		reader := bufio.NewReader(f)

		_, err = io.Copy(w, reader)
		checkError(err)

		err = os.Remove(name)
		checkError(err)

		w.Flush()
	}
}
