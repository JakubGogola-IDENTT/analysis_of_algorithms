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
func WithNodes(nodesCount int) (slotsCount int) {
	slot := NONE

	for slot != SINGLE {
		slot = NONE
		slotsCount++

		for i := 0; i < nodesCount; i++ {
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
func WithUpperLimit(upperLimit, nodesCount int) (slotsCount int, roundsCount int) {
	slot := NONE

	limit := int(math.Ceil(math.Log2(float64(upperLimit))))

	for slot != SINGLE {
		roundsCount++
		for i := 1; i <= limit; i++ {
			slot = NONE
			slotsCount++

			for j := 0; j < nodesCount; j++ {
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

	return slotsCount, roundsCount
}
