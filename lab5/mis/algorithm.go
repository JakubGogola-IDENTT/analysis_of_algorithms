package mis

type Graph struct {
	vs []vertex
	es []edge
}

type vertex struct {
	id   int
	pref *vertex
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

func (g *Graph) isMarried(p *vertex) bool {
	pAdj := g.getAdj(p)

	q := p.pref
	qAdj := g.getAdj(q)

	return p.pref.id == q.pref.id && contains(pAdj, p) && contains(qAdj, p) // TODO: check if it works
}

func (g *Graph) isSingle(p *vertex) bool {
	adj := g.getAdj(p)

	allMarried := true
	for _, a := range adj {
		allMarried = allMarried && g.isMarried(a)
	}

	return p.pref == nil && allMarried
}

func (g *Graph) isStable() bool {
	for i := range g.vs {
		if !g.isSingle(&g.vs[i]) || g.isMarried(&g.vs[i]) {
			return false
		}
	}

	return true
}

func (g *Graph) process(p *vertex) {
	var q *vertex
	adj := g.getAdj(p)

	// accept proposal
	if p.pref == nil {

		q = nil

		for i, a := range adj {
			if a.pref != nil && a.pref.id == p.id {
				q = adj[i]
				break
			}
		}

		if q != nil {
			p.pref = q
		}
	}

	// propose
	allNotMarriedWithP := true
	someFree := false
	q = nil

	for i, a := range adj {
		if a.pref == nil {
			someFree = true
			q = adj[i]
			continue
		}

		if a.pref.id == p.id {
			allNotMarriedWithP = false
			break
		}
	}

	if p.pref == nil && allNotMarriedWithP && someFree {
		p.pref = q
	}

	// unchain
	q = p.pref

	if q != nil && q.pref != nil && q.pref.id != p.id {
		p.pref = nil
	}
}

func New(n int) (g Graph) {
	g.vs = make([]vertex, n)

	for i := range g.vs {
		g.vs[i] = vertex{
			id: i,
		}
	}

	for i, p := range g.vs {
		for j, q := range g.vs {
			if p.id < q.id {
				g.es = append(g.es, edge{
					v1: &g.vs[i],
					v2: &g.vs[j],
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

func (g *Graph) Simulate() {
	isStable := false
	for !isStable {
		for i := range g.vs {
			g.process(&g.vs[i])

			if g.isStable() {
				break
			}
		}
	}
}
