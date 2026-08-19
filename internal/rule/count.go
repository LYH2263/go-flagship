package rule

func CountByOp(specs []Spec) map[string]int {
	out := make(map[string]int)
	for _, s := range specs {
		out[NormalizeOp(s.Op)]++
	}
	return out
}
