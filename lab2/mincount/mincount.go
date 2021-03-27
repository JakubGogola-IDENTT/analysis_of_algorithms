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

func New(hash func() hash.Hash, k int) MinCount {
	return MinCount{
		Hash: hash(),
		K:    k,
	}
}

func (mc *MinCount) Sum(multiset []int) int {
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
