package main

import (
	"crypto/sha256"
	"fmt"
	"lab3/hyperloglog"
)

func main() {
	fmt.Println("Hello, world!")

	hll := hyperloglog.New(sha256.New, 4, 10)

	for i := 0; i < 100; i++ {
		hll.Add(i)
	}
}
