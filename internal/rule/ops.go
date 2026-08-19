package rule

func KnownOps() []string {
	return []string{"eq", "ne", "in", "notin", "exists", "missing"}
}

func NormalizeOp(op string) string {
	switch op {
	case "EQ", "Eq":
		return "eq"
	case "NE", "Ne":
		return "ne"
	case "IN", "In":
		return "in"
	case "NOTIN", "NotIn", "notin":
		return "notin"
	case "EXISTS", "Exists":
		return "exists"
	case "MISSING", "Missing":
		return "missing"
	default:
		return op
	}
}
