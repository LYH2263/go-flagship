package flagship

// CloneEvalAttrs 拷贝评估属性。
func CloneEvalAttrs(a EvalAttrs) EvalAttrs {
	if a == nil {
		return nil
	}
	out := make(EvalAttrs, len(a))
	for k, v := range a {
		out[k] = v
	}
	return out
}
