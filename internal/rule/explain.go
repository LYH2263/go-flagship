package rule

import "strings"

func Explain(s Spec) string {
	var b strings.Builder
	b.WriteString(s.Attr)
	b.WriteByte(' ')
	b.WriteString(s.Op)
	if len(s.Values) > 0 {
		b.WriteByte(' ')
		b.WriteString(strings.Join(s.Values, ","))
	}
	return b.String()
}

func ExplainAll(specs []Spec) string {
	parts := make([]string, 0, len(specs))
	for _, s := range specs {
		parts = append(parts, Explain(s))
	}
	return strings.Join(parts, " AND ")
}
