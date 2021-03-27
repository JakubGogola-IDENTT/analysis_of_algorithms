package mincount

import (
	"encoding/binary"
)

func (mc *MinCount) getHash(v int) float64 {
	converted := string(rune(v))

	hash := mc.Hash.Sum([]byte(converted))

	val := binary.BigEndian.Uint64(hash)

	return float64(val>>11) / float64(1<<53)
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
