package election

import (
	"math"
	"math/rand"
)

const (
	none      string = "none"
	single    string = "single"
	collision string = "collision"
)

// ElectByScenario2 choses leader using second scenario from lecture (exact number of nodes is given)
func ElectByScenario2(nodesCount int) (slotsCount int) {
	broadcastProbability := 1.0 / float64(nodesCount)
	slot := none

	for slot != single {
		slot = none
		slotsCount++

		for i := 0; i < nodesCount; i++ {
			if rand.Float64() >= broadcastProbability {
				continue
			}

			if slot == none {
				slot = single
			} else if slot == single {
				slot = collision
			}
		}
	}

	return slotsCount
}

// ElectByScenario3 choses leader using third scenario from lecture (only upper limit for nodes is given)
func ElectByScenario3(upperLimit, nodesCount int) (slotsCount, roundsCount int) {
	slot := none
	limit := int(math.Ceil(math.Log2(float64(upperLimit))))

	for slot != single {
		roundsCount++
		slotsCount = 0

		for i := 1; i <= limit; i++ {
			slot = none
			slotsCount++

			for j := 0; j < nodesCount; j++ {
				if rand.Float64() >= 1.0/math.Exp2(float64(i)) {
					continue
				}

				if slot == none {
					slot = single
				} else if slot == single {
					slot = collision
				}
			}

			if slot == single {
				break
			}
		}
	}

	return slotsCount, roundsCount
}
