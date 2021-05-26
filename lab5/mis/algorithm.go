package mis

type Graph struct {
	vs    []int
	es    []edge
	prefs map[int]int
}

type vertex struct{
	id int
	pref *vertex
}

type edge struct {
	i, j int
}

func (g *Graph) getAdj(v int) (adj []int) {
	if v == -1 {
		return adj
	}

	for _, e := range g.es {
		if v == e.i {
			adj = append(adj, e.j)
		}

		if v == e.j {
			adj = append(adj, e.i)
		}
	}

	return adj
}

func (g *Graph) isMarried(p int) bool {
	pAdj := g.getAdj(p)

	q := g.prefs[p]
	qAdj := g.getAdj(q)

	return contains(p, pAdj) && contains(g.prefs[q], qAdj) // TODO: check if it works
}

func (g *Graph) process(p int) {
	adj := g.getAdj(p)

	if g.prefs[p] == -1 {
		q := -1
		for _, a := range adj {
			if next := g.prefs[a]; next == -1 {
				q = next
				break
			}
		}

		if q != -1 {
			g.prefs[p] = q
		}
	}

	if g.prefs[p] == -1 {
		anyFree := false
		for _, a := range adj {
			if g.prefs[a] 
		}
	}
}

func New(n int) (g Graph) {
	g.vs = make([]int, n)
	g.prefs = make(map[int]int, n)

	for i := range g.vs {
		g.vs[i] = i
	}

	for _, i := range g.vs {
		for _, j := range g.vs {
			if i < j {
				g.es = append(g.es, edge{i, j})
			}
		}
	}

	for idx, e := range g.es {
		if e.i > e.j {
			g.es[idx].i, g.es[idx].j = g.es[idx].j, g.es[idx].i
		}
	}

	return g
}

func (g *Graph) Simulate() {
	// TODO: clear graph prefs

	for k := range g.prefs {
		g.prefs[k] = -1
	}

	isStable := false
	for !isStable {
		for _, v := range g.vs {
			g.process(v)
		}
	}
}
