package main

import (
	"crypto/sha256"
	"fmt"
	"lab3/hyperloglog"
)

func main() {
	hll := hyperloglog.New(sha256.New, 16)

	for i := 0; i < 100000; i++ {
		hll.Add(i)
	}

	est := hll.Count()

	fmt.Println(est)
}
