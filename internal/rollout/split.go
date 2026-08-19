package rollout

// SplitBuckets 将 0..99 按百分比切分。
func SplitBuckets(percent int) (onMax, offMin int) {
	p := ClampPercent(percent)
	return p, p
}
