package hyperloglog

import (
	"hash"
	"log"
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

func (hll *HyperLogLog) Count() {
	estimate := hll.getEstimate()

}
