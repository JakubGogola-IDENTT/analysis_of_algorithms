package hyperloglog

import (
	"hash"
	"log"
	"math"
)

type HyperLogLog struct {
	Hash hash.Hash
	M    int
	B    int
	regs []uint8
}

func New(hashFunc func() hash.Hash, m, b int) HyperLogLog {
	if b < 4 || b > 16 {
		log.Fatal("b should be in [4;16]")
	}

	return HyperLogLog{
		Hash: hashFunc(),
		M:    m,
		B:    b,
		regs: make([]uint8, m),
	}
}

func (hll *HyperLogLog) Add(value int) {
	hash := hll.getHash(value)

	j := hll.eb32(hash)
	w := hash<<hll.B | j<<(hll.B-1)

	zeroBits := hll.clz32(w) + 1

	if zeroBits > hll.regs[j] {
		hll.regs[j] = zeroBits
	}
}

func (hll *HyperLogLog) Count() uint64 {
	estimate := hll.getEstimate()

	m := float64(hll.M)
	if estimate <= m*2.5 {
		if c := hll.countZeroes(); c > 0 {
			return uint64(m * math.Log(m/float64(c)))
		}

		return uint64(estimate)
	} else if estimate < float64(twoTo32)/30. {
		return uint64(estimate)
	}

	return uint64(-float64(twoTo32) * math.Log(1.-estimate/float64(twoTo32)))
}
