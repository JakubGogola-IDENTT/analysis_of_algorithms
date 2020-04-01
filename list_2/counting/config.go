package counting

import (
	"flag"
	"fmt"
)

// MinCount is struct which holds data for algorithm
type MinCount struct {
	multisetSize    int
	k               int
	testMode        bool
	withRepetitions bool
}

func (mc *MinCount) parseFlags() {
	flag.IntVar(&mc.k, "k", 2, "size of array with hashes (default: 2)")
	flag.IntVar(&mc.multisetSize, "multisetSize", 1000, "size of multiset (default: 100)")
	flag.BoolVar(&mc.testMode, "testMode", false, "test mode")
	flag.BoolVar(&mc.withRepetitions, "withRepetitions", false, "with repetitions")
	flag.Parse()
}

// Init initializes algorithm with given parameters
func (mc *MinCount) Init() {
	mc.parseFlags()

	multiset := generateMultisetWithRepetitions(mc.multisetSize)

	if mc.testMode {
		runTests(mc.withRepetitions)
		return
	}

	count := counting(multiset, mc.k)

	fmt.Println(count)
}
