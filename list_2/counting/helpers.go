package counting

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/rand"
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

	hashFunc := md5.New()
	hashFunc.Write(data)

	hashBytes := hashFunc.Sum(nil)

	hash = float64(binary.BigEndian.Uint32(hashBytes)) / float64(1<<32-1)

	return hash
}

func countNonOnes(arr []float64) (count int) {
	for _, elt := range arr {
		if elt != 1. {
			count++
		}
	}

	return count
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

func generateMultisetWithRepetitions(size int) []int {
	var multiset []int

	for i := 0; i < size; i++ {
		repetitions := rand.Intn(5)

		for j := 0; j < repetitions; j++ {
			multiset = append(multiset, i)
		}
	}

	return multiset
}

func getFileName(k int, withRepetitions bool) string {
	if withRepetitions {
		return fmt.Sprintf("k_%d_rep.csv", k)
	}

	return fmt.Sprintf("k_%d.csv", k)
}
