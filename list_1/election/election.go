package election

import (
	"math"
	"math/rand"
)

const (
	// NONE - no slot was elected
	NONE = "none"
	// SINGLE - single slot was elected
	SINGLE = "single"
	// COLLISION - more than one slot was elected
	COLLISION = "collision"
)

// WithNodes - simulates leader election with given number of nodes
func WithNodes(nodesCount uint) uint {
	slot := NONE
	slotsCount := uint(0)

	for slot != SINGLE {
		slot = NONE
		slotsCount++

		for i := uint(0); i < nodesCount; i++ {
			broadcastProb := rand.Float64()

			if broadcastProb >= 1.0/float64(nodesCount) {
				continue
			}

			if slot == NONE {
				slot = SINGLE
			} else if slot == SINGLE {
				slot = COLLISION
			}
		}
	}

	return slotsCount
}

// WithUpperLimit - simulates lider election with given upper limit of nodes
func WithUpperLimit(upperLimit uint, nodesCount uint) (uint, uint) {
	slot := NONE
	slotsCount := uint(0)
	roundsCount := uint(0)

	limit := uint(math.Ceil(math.Log2(float64(upperLimit))))

	for slot != SINGLE {
		roundsCount++
		for i := uint(1); i <= limit; i++ {
			slot = NONE
			slotsCount++

			for j := uint(0); j < nodesCount; j++ {
				broadcastProb := rand.Float64()

				if broadcastProb >= 1.0/math.Exp2(float64(i)) {
					continue
				}

				if slot == NONE {
					slot = SINGLE
				} else if slot == SINGLE {
					slot = COLLISION
				}
			}

			if slot == SINGLE {
				break
			}
		}
	}

	return roundsCount, slotsCount
}
