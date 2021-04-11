package hyperloglog

import (
	"hash"
	"log"
	"math"
)

type HyperLogLog struct {
	hash hash.Hash
	m    int
	b    int
	regs []uint8
}

func New(hashFunc func() hash.Hash, b int) HyperLogLog {
	if b < 4 || b > 16 {
		log.Fatal("b should be in [4;16]")
	}

	m := 1 << b

	return HyperLogLog{
		hash: hashFunc(),
		m:    m,
		b:    b,
		regs: make([]uint8, m),
	}
}

func (hll *HyperLogLog) Clear() {
	hll.regs = make([]uint8, hll.m)
}

func (hll *HyperLogLog) Add(value int) {
	hash := hll.getHash(value)

	j := hll.eb32(hash)
	w := hash<<hll.b | j<<(hll.b-1)

	zeroBits := hll.clz32(w) + 1

	if zeroBits > hll.regs[j] {
		hll.regs[j] = zeroBits
	}
}

func (hll *HyperLogLog) Count() uint64 {
	estimation := hll.getEstimate()

	m := float64(hll.m)
	if estimation <= m*2.5 {
		if c := hll.countZeroes(); c != 0 {
			return uint64(m * math.Log(m/float64(c)))
		}

		return uint64(estimation)
	} else if estimation < twoTo32/30. {
		return uint64(estimation)
	}

	return uint64(-twoTo32 * math.Log(1.-estimation/twoTo32))
}
