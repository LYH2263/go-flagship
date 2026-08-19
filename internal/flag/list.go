package flag

import "sort"

func Keys(t *Table) []string {
	if t == nil {
		return nil
	}
	out := make([]string, 0, t.Len())
	for _, r := range t.Export() {
		out = append(out, r.Key)
	}
	sort.Strings(out)
	return out
}
