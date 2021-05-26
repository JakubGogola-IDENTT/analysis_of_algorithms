package mis

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Graph struct {
	vs []vertex
	es []edge
}

type vertex struct {
	id  int
	mis bool
}

type edge struct {
	v1, v2 *vertex
}

func (g *Graph) getAdj(v *vertex) (adj []*vertex) {
	if v == nil {
		return adj
	}

	for i, e := range g.es {
		if v.id == e.v1.id {
			adj = append(adj, g.es[i].v2)
		}

		if v.id == e.v2.id {
			adj = append(adj, g.es[i].v1)
		}
	}

	return adj
}

func (g *Graph) isSomeAdjInMIS(p *vertex) bool {
	adj := g.getAdj(p)

	for _, a := range adj {
		if a.mis {
			return true
		}
	}

	return false
}

func (g *Graph) isConf(p *vertex) bool {
	return p.mis && g.isSomeAdjInMIS(p)
}

func (g *Graph) isCand(p *vertex) bool {
	adj := g.getAdj(p)
	allAdjNotInMis := true

	for _, a := range adj {
		if a.mis {
			allAdjNotInMis = false
			break
		}
	}

	return !p.mis && allAdjNotInMis
}

func (g *Graph) isAck(p *vertex) bool {
	return p.mis && !g.isSomeAdjInMIS(p)
}

func (g *Graph) isNAck(p *vertex) bool {
	return !p.mis && g.isSomeAdjInMIS(p)
}

func (g *Graph) isStable() bool {
	allStable := true

	for i := range g.vs {
		v := &g.vs[i]

		allStable = g.isAck(v) || g.isNAck(v)
	}

	return allStable
}

func (g *Graph) process(p *vertex) {
	if g.isConf(p) {
		p.mis = false
	}

	if g.isCand(p) {
		p.mis = true
	}
}

func New(n int) (g Graph) {
	g.vs = make([]vertex, n)

	for i := range g.vs {
		g.vs[i] = vertex{
			id:  i,
			mis: false,
		}
	}

	for i := range g.vs {
		p := g.vs[i]
		for j := range g.vs {
			q := g.vs[j]

			if p.id < q.id {
				g.es = append(g.es, edge{
					v1: &p,
					v2: &q,
				})
			}
		}
	}

	for idx, e := range g.es {
		if e.v1.id > e.v2.id {
			g.es[idx].v1, g.es[idx].v2 = g.es[idx].v2, g.es[idx].v1
		}
	}

	return g
}

func (g *Graph) PrintMIS() {
	for i := range g.vs {
		v := &g.vs[i]
		adj := g.getAdj(v)

		var adjLabels []string

		for _, a := range adj {
			adjLabels = append(adjLabels, fmt.Sprintf("%d: %t", a.id, a.mis))
		}

		fmt.Printf("%d: %t -> %s\n", v.id, v.mis, strings.Join(adjLabels, ", "))
	}
}

func (g *Graph) Simulate() {
	rand.Seed(time.Now().UnixNano())

	for {
		for i := range g.vs {
			v := &g.vs[i]
			g.process(v)
		}

		if g.isStable() {
			break
		}
	}
}
