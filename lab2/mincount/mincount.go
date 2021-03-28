package mincount

import (
	"hash"
	"sort"
)

// MinCount contains all required data for MinCount algorithm
type MinCount struct {
	Hash        hash.Hash
	K           int
	HashBitsLen int
}

func NewWithHashBitsLen(hashFunc func() hash.Hash, k, hashBitsLen int) MinCount {
	return MinCount{
		Hash:        hashFunc(),
		K:           k,
		HashBitsLen: hashBitsLen,
	}
}

func New(hashFunc func() hash.Hash, k int) MinCount {
	return NewWithHashBitsLen(hashFunc, k, hashFunc().Size()*8)
}

func (mc *MinCount) Count(multiset []int) int {
	hashes := mc.hashesList()

	for _, x := range multiset {
		hash := mc.getHash(x)

		if hash < hashes[mc.K-1] && !contains(hash, hashes) {
			hashes[mc.K-1] = hash

			sort.Float64s(hashes)
		}
	}

	if hashes[mc.K-1] == 1. {
		return countNotEqual(1., hashes)
	}

	return int(float64(mc.K-1) / hashes[mc.K-1])
}
