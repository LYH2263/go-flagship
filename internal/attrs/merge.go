package attrs

func Merge(base, overlay map[string]string) map[string]string {
	out := CloneMap(base)
	if out == nil {
		out = make(map[string]string)
	}
	for k, v := range overlay {
		out[k] = v
	}
	return out
}
