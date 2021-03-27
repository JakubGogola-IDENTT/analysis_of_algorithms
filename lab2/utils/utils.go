package utils

import (
	"crypto/rand"
	"log"
	"math/big"
)

// GetRandomInt returns random integer from given range using crypto/rand package.
func GetRandomInt(size int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(size)))

	if err != nil {
		log.Fatal(err)
	}

	return int(n.Int64())
}

func CountDistinct(vs []int) int {
	diffVs := make(map[int]bool)

	for _, v := range vs {
		diffVs[v] = true
	}

	return len(diffVs)
}
