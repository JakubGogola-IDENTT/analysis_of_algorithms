package anlaytics

import (
	"crypto/rand"
	"lab2/utils"
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
	distinctVs := make(map[int]bool)

	for _, v := range vs {
		distinctVs[v] = true
	}

	return len(distinctVs)
}

func createMultiset(size, randRange int) (multiset []int) {
	multiset = make([]int, size)

	for i := 0; i < size; i++ {
		multiset[i] = utils.GetRandomInt(randRange)
	}

	return multiset
}
