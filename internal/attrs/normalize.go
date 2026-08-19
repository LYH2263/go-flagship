package attrs

import "strings"

func NormalizeKey(k string) string {
	return strings.TrimSpace(strings.ToLower(k))
}

func NormalizeMap(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	out := make(map[string]string, len(m))
	for k, v := range m {
		out[NormalizeKey(k)] = strings.TrimSpace(v)
	}
	return out
}
