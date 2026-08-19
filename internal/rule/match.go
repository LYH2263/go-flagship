package rule

import "strings"

func MatchAll(specs []Spec, attrs map[string]string) bool {
	for _, s := range specs {
		if !MatchOne(s, attrs) {
			return false
		}
	}
	return true
}

func MatchOne(s Spec, attrs map[string]string) bool {
	op := strings.ToLower(s.Op)
	val, ok := "", false
	if attrs != nil {
		val, ok = attrs[s.Attr]
	}
	switch op {
	case "exists":
		return ok
	case "missing":
		return !ok
	case "eq":
		if !ok || len(s.Values) == 0 {
			return false
		}
		return val == s.Values[0]
	case "ne":
		if !ok || len(s.Values) == 0 {
			return true
		}
		return val != s.Values[0]
	case "in":
		if !ok {
			return false
		}
		for _, v := range s.Values {
			if val == v {
				return true
			}
		}
		return false
	case "notin":
		if !ok {
			return true
		}
		for _, v := range s.Values {
			if val == v {
				return false
			}
		}
		return true
	default:
		return false
	}
}
