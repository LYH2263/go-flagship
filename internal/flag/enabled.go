package flag

func EnabledKeys(t *Table) []string {
	if t == nil {
		return nil
	}
	var out []string
	for _, r := range t.Export() {
		if r.Enabled {
			out = append(out, r.Key)
		}
	}
	return out
}
