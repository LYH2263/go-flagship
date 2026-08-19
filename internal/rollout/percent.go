package rollout

func ClampPercent(p int) int {
	if p < 0 {
		return 0
	}
	if p > 100 {
		return 100
	}
	return p
}

func Describe(percent int) string {
	p := ClampPercent(percent)
	switch {
	case p == 0:
		return "none"
	case p == 100:
		return "all"
	default:
		return "partial"
	}
}
