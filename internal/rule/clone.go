package rule

func CloneSpecs(in []Spec) []Spec {
	if in == nil {
		return nil
	}
	out := make([]Spec, len(in))
	for i, s := range in {
		out[i] = Spec{
			Attr:   s.Attr,
			Op:     s.Op,
			Values: append([]string(nil), s.Values...),
		}
	}
	return out
}

func CloneValues(in []string) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	copy(out, in)
	return out
}
