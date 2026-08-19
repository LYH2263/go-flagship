package eval

func RunAll(inputs []Input) []Result {
	out := make([]Result, 0, len(inputs))
	for _, in := range inputs {
		out = append(out, Run(in))
	}
	return out
}
