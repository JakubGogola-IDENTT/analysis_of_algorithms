package analysis

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func printProgress(n int) {
	if n%100 == 0 {
		fmt.Printf("Progress: %d\n", n)
	}
}
