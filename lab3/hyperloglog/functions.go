package hyperloglog

import (
	"io"
	"math/big"
	"strconv"
)

const (
	alpha16 float64 = 0.673
	alpha32 float64 = 0.697
	alpha64 float64 = 0.709
)

func (hll *HyperLogLog) getAlpha() float64 {
	switch hll.M {
	case 16:
		return alpha16
	case 32:
		return alpha32
	case 64:
		return alpha64
	default:
		return 0.7213 / (1 + 1.079/float64(hll.M))
	}
}

func (hll *HyperLogLog) getHash(v int) uint32 {
	var hashVal big.Int

	hll.Hash.Reset()
	io.WriteString(hll.Hash, strconv.Itoa(v))
	hash := hll.Hash.Sum(nil)

	hashVal.SetBytes(hash)

	if hashLen := hll.Hash.Size() * 8; hashLen > 32 {
		shift := hashLen - 32
		hashVal.Rsh(&hashVal, uint(shift))
	}

	return uint32(hashVal.Int64())
}

func (hll *HyperLogLog) eb32(v uint32) uint32 {
	m := uint32(((1 << (32 - hll.B)) - 1) << hll.B)
	return (v & m) >> hll.B
}

func (hll *HyperLogLog) getEstimate() float64 {
	sum := 0.

	for _, r := range hll.regs {
		sum += 1.0 / float64(uint64(1)<<r)
	}

	m := float64(hll.M)
	return hll.getAlpha() * m * m / sum
}

// https://embeddedgurus.com/state-space/2014/09/fast-deterministic-and-portable-counting-leading-zeros/
// Optimal counting of leading zeroes
func (hll *HyperLogLog) clz32(v uint32) uint8 {
	var clzLookup = []uint8{
		32, 31, 30, 30, 29, 29, 29, 29, 28, 28, 28, 28, 28, 28, 28, 28,
	}

	var n uint8

	if v >= (1 << 16) {
		if v >= (1 << 24) {
			if v >= (1 << 28) {
				n = 28
			} else {
				n = 24
			}
		} else {
			if v >= (1 << 20) {
				n = 20
			} else {
				n = 16
			}
		}
	} else {
		if v >= (1 << 8) {
			if v >= (1 << 12) {
				n = 12
			} else {
				n = 8
			}
		} else {
			if v >= (1 << 4) {
				n = 4
			} else {
				n = 0
			}
		}
	}
	return clzLookup[v>>n] - n
}
