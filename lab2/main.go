package main

import (
	"crypto/sha512"
	"fmt"
	mc "lab2/mincount"
	"lab2/utils"
)

func main() {
	multiset := make([]int, 1000)

	for i := 0; i < len(multiset); i++ {
		multiset[i] = utils.GetRandomInt(1000)
	}

	// fmt.Println(multiset)
	fmt.Printf("Expected: %d\n", utils.CountDistinct(multiset))

	algorithm := mc.New(sha512.New, len(multiset))
	n := algorithm.Sum(multiset)

	fmt.Printf("Computed: %d\n", n)
}
