package mis

func contains(vs []*vertex, p *vertex) bool {
	for _, v := range vs {
		if v.id == p.id {
			return true
		}
	}

	return false
}
