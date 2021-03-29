package mincount

import (
	"io"
	"math/big"
	"strconv"
	"strings"
)

func (mc *MinCount) XD(v int) float64 {
	return mc.getHash(v)
}

func (mc *MinCount) getHash(v int) float64 {
	var hashVal, maxBinVal big.Int

	mc.Hash.Reset()
	io.WriteString(mc.Hash, strconv.Itoa(v))
	hash := mc.Hash.Sum(nil)

	hashVal.SetBytes(hash)

	hashLen := mc.HashBitsLen
	maxHashLen := mc.Hash.Size() * 8

	if maxHashLen < mc.HashBitsLen {
		hashLen = maxHashLen
	}

	divideBy := len(hashVal.Bytes())*8 - hashLen

	hashVal.Rsh(&hashVal, uint(divideBy))

	maxBinVal.SetString(strings.Repeat("1", mc.HashBitsLen), 2)

	return float64(hashVal.Uint64()) / float64(maxBinVal.Uint64())
}

func (mc *MinCount) hashesList() (hashes []float64) {
	hashes = make([]float64, mc.K)

	for i := 0; i < mc.K; i++ {
		hashes[i] = 1.
	}

	return hashes
}

func contains(vs []float64, n float64) bool {
	for _, v := range vs {
		if v == n {
			return true
		}
	}

	return false
}

func countNotEqual(n float64, vs []float64) (count int) {
	for _, v := range vs {
		if v != n {
			count++
		}
	}

	return count
}
