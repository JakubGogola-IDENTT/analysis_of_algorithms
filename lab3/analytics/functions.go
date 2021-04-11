package analytics

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func (a *Analytics) parseFlags() {
	flag.IntVar(&a.minSize, "l", 1, "Minimal size of set")
	flag.IntVar(&a.maxSize, "u", 10000, "Maximal size of set")
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

func printProgress(hashFuncName string, b, n, div int) {
	if n%div == 0 {
		fmt.Printf("%s | b = %d | Progress: %d...\n", hashFuncName, b, n)
	}
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

func createMultiset(start, size int) (multiset []int, next int) {
	multiset = make([]int, size)
	next = start + size // start - inclusive, next - exlusive

	for i := 0; i < size; i++ {
		multiset[i] = i + start
	}

	return multiset, next
}
