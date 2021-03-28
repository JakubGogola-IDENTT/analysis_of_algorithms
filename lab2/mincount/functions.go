package mincount

import (
	"fmt"
	"io"
	"math/big"
	"strconv"
)

func (mc *MinCount) getHash(v, bits int) float64 {
	converted := strconv.Itoa(v)

	mc.Hash.Reset()
	io.WriteString(mc.Hash, converted)
	hash := mc.Hash.Sum(nil)

	var value, hashVal big.Int

	hashVal.SetBytes(hash)
	fmt.Println(hashVal.String())

	// Lenght of binary representation of hash - bits limit
	divider := uint(mc.Hash.Size()*8 - bits)
	value.Rsh(&hashVal, divider)

	maxVal := (1 << bits) - 1

	return float64(value.Int64()) / float64(maxVal)
}

func (mc *MinCount) XD(v, bits int) float64 {
	return mc.getHash(v, bits)
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
