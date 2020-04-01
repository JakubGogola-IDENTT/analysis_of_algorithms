package counting

import (
	"flag"
	"fmt"
)

// MinCount is struct which holds data for algorithm
type MinCount struct {
	multisetSize int
	k            int
	testMode     bool
}

func (mc *MinCount) parseFlags() {
	flag.IntVar(&mc.k, "k", 300, "size of array with hashes (default: 2)")
	flag.IntVar(&mc.multisetSize, "multisetSize", 1000, "size of multiset (default: 100)")
	flag.BoolVar(&mc.testMode, "testMode", false, "test mode")
	flag.Parse()
}

// Init initializes algorithm with given parameters
func (mc *MinCount) Init() {
	mc.parseFlags()

	multiset := generateMultiset(mc.multisetSize)
	count := counting(multiset, mc.k)

	fmt.Println(count)
}
