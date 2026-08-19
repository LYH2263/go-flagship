package attrs

func FilterKeys(m map[string]string, keep []string) map[string]string {
	if m == nil {
		return nil
	}
	out := make(map[string]string)
	set := make(map[string]struct{}, len(keep))
	for _, k := range keep {
		set[k] = struct{}{}
	}
	for k, v := range m {
		if _, ok := set[k]; ok {
			out[k] = v
		}
	}
	return out
}
