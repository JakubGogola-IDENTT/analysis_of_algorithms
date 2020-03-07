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

	fmt.Printf("#################################\n")
	fmt.Printf("With nodes (n = %d)\n", nodesCount)
	fmt.Printf("#################################\n")
	computeStats(data, iterationsCount)
}

func withUpperLimitAnalysis(upperLimit, nodesCount, iterationsCount int) {
	data := make(map[int]int)

	for i := 0; i < iterationsCount; i++ {
		slot, _ := election.WithUpperLimit(upperLimit, nodesCount)

		data[slot]++
	}

	fmt.Printf("#################################\n")
	fmt.Printf("With upper limit (u = %d, n = %d)\n", upperLimit, nodesCount)
	fmt.Printf("#################################\n")
	createHistogram(data, fmt.Sprintf("with_upper_limit_u_%d_n_%d", upperLimit, nodesCount))
	computeStats(data, iterationsCount)
	computerFirstRoundSuccessProbability(upperLimit, nodesCount, iterationsCount)
}
