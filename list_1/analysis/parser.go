package analysis

import (
	"flag"
	"log"
)

type arguments struct {
	n int
	u int
	i int
}

func (a *arguments) parseArgs() {
	flag.IntVar(&a.i, "i", 1000, "number of iterations")
	flag.IntVar(&a.u, "u", 1000, "upper limit for nodes")
	flag.IntVar(&a.n, "n", 1000, "number of nodes")

	if a.u < 2 {
		log.Fatal("'u' argument should be greater than 2")
	}
}
