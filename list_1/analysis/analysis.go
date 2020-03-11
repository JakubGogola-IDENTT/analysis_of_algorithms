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

	if args.t {
		runTests(args)
		return
	}

	withNodesAnalysis(args.n, args.i, true, true)

	withUpperLimitAnalysis(args.u, 2, args.i, true, true)
	withUpperLimitAnalysis(args.u, args.u/2, args.i, true, true)
	withUpperLimitAnalysis(args.u, args.u, args.i, true, true)
}

func withNodesAnalysis(nodesCount, iterationsCount int, withStats bool, withHistogram bool) {
	data := make(map[int]int)

	for i := 0; i < iterationsCount; i++ {
		slot := election.WithNodes(nodesCount)

		data[slot]++
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("With nodes (n = %d)\n", nodesCount)
	fmt.Println("-------------------------------------")

	if withStats {
		computeStats(data, iterationsCount)
	}

	if withHistogram {
		createHistogram(data, fmt.Sprintf("with_nodes_n_%d", nodesCount))
	}
}

func withUpperLimitAnalysis(upperLimit, nodesCount, iterationsCount int, withStats bool, withHistogram bool) {
	data := make(map[int]int)

	for i := 0; i < iterationsCount; i++ {
		slot, _ := election.WithUpperLimit(upperLimit, nodesCount)

		data[slot]++
	}

	fmt.Println("-------------------------------------")
	fmt.Printf("With upper limit (u = %d, n = %d)\n", upperLimit, nodesCount)
	fmt.Println("-------------------------------------")

	if withStats {
		computerFirstRoundSuccessProbability(upperLimit, nodesCount, iterationsCount)
	}

	if withHistogram {
		createHistogram(data, fmt.Sprintf("with_upper_limit_u_%d_n_%d", upperLimit, nodesCount))
	}
}

func runTests(args arguments) {
	// task 2
	fmt.Println("Task 2")
	withNodesAnalysis(args.n, args.i, true, true)
	withUpperLimitAnalysis(args.u, 2, args.i, false, true)
	withUpperLimitAnalysis(args.u, args.u/2, args.i, false, true)
	withUpperLimitAnalysis(args.u, args.u, args.i, false, true)

	// task 3
	fmt.Println("Task 3")
	for n := 1; n <= args.n; n++ {
		withNodesAnalysis(n, args.i, true, false)
	}

	// task 4
	fmt.Println("Task 4")
	withUpperLimitAnalysis(args.u, 2, args.i, true, false)
	withUpperLimitAnalysis(args.u, args.u/2, args.i, true, false)
	withUpperLimitAnalysis(args.u, args.u, args.i, true, false)

}
