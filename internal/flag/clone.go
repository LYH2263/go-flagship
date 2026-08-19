package flag

func CloneRules(in []Rule) []Rule {
	if in == nil {
		return nil
	}
	out := make([]Rule, len(in))
	for i, r := range in {
		out[i] = Rule{
			Attr:   r.Attr,
			Op:     r.Op,
			Values: append([]string(nil), r.Values...),
		}
	}
	return out
}
