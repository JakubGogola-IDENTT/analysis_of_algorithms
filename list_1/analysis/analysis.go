package analysis

import (
	"fmt"
	"list_1/election"
	"math/rand"
	"time"
)

// Run - runs tests
func Run() {
	rand.Seed(time.Now().UnixNano())

	args := arguments{}
	args.parseArgs()

	withNodesAnalysis(args.n, args.i)

	withUpperLimitAnalysis(args.u, 2, args.i)
	withUpperLimitAnalysis(args.u, args.u/2, args.i)
	withUpperLimitAnalysis(args.u, args.u, args.i)
}

func withNodesAnalysis(nodesCount, iterationsCount int) {
	data := make(map[int]int)

	for i := 0; i < iterationsCount; i++ {
		slot := election.WithNodes(nodesCount)

		data[slot]++
	}

	createHistogram(data, "with_nodes")
	computeStats(data, iterationsCount)
}

func withUpperLimitAnalysis(upperLimit, nodesCount, iterationsCount int) {
	data := make(map[int]int)

	for i := 0; i < iterationsCount; i++ {
		slot, _ := election.WithUpperLimit(upperLimit, nodesCount)

		data[slot]++
	}

	createHistogram(data, fmt.Sprintf("with_upper_limit_u_%d_n_%d", upperLimit, nodesCount))
}
