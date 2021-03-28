package mincount

import (
	"io"
	"math/big"
	"strconv"
	"strings"
)

func (mc *MinCount) getHash(v int) float64 {
	// TODO: probably requires deep refactor
	converted := strconv.Itoa(v)

	mc.Hash.Reset()
	io.WriteString(mc.Hash, converted)
	hash := mc.Hash.Sum(nil)

	var value, hashVal, maxBinVal big.Int

	hashVal.SetBytes(hash)

	length := mc.HashBitsLen
	maxLen := mc.Hash.Size() * 8

	if maxLen < mc.HashBitsLen {
		length = maxLen
	}

	// Lenght of binary representation of hash - bits limit
	divider := uint(hashVal.BitLen() - length)

	value.Rsh(&hashVal, divider)

	maxBinRepr := strings.Repeat("1", length)
	maxBinVal.SetString(maxBinRepr, 2)

	return float64(value.Uint64()) / float64(maxBinVal.Uint64())
}

func (mc *MinCount) hashesList() (hashes []float64) {
	hashes = make([]float64, mc.K)

	for i := 0; i < mc.K; i++ {
		hashes[i] = 1.
	}

	return hashes
}

func contains(n float64, vs []float64) bool {
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
