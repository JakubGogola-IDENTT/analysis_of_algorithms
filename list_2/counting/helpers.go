package counting

import (
	"crypto/sha256"
	"encoding/binary"
	"strconv"
)

func ones(k int) []float64 {
	arr := make([]float64, k)

	for i := range arr {
		arr[i] = 1.
	}

	return arr
}

func getHash(value int) (hash float64) {
	data := []byte(strconv.Itoa(value))

	hashFunc := sha256.New()
	hashFunc.Write(data)

	hashBytes := hashFunc.Sum(nil)

	hash = float64(binary.BigEndian.Uint32(hashBytes)) / float64(1<<32-1)

	return hash
}

func includes(value float64, arr []float64) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func generateMultiset(size int) []int {
	var multiset = make([]int, size)

	for i := 0; i < size; i++ {
		multiset[i] = i
	}

	return multiset
}
