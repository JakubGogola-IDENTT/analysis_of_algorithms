package analysis

import "flag"

// Analysis struct contains all params required for analysis of algorithm
type Analysis struct {
	nodes    int
	limit    int
	scenario int
}

func (a *Analysis) parseFlags() {
	// TODO: update default values
	flag.IntVar(&a.nodes, "nodes", 100, "number of nodes")
	flag.IntVar(&a.limit, "limit", 100, "limit used for third scenario")
	flag.IntVar(&a.scenario, "scenario", 2, "scenario number - 2 or 3")
	flag.Parse()
}

// TestScenario2 tests second scenario implementation
func (a *Analysis) TestScenario2() {

}
