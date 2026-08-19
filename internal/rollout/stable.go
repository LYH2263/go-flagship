package rollout

// StableOn 对同一 subject 稳定落在百分比内。
func StableOn(flagKey, subject string, percent int) bool {
	_, on := Decide(DefaultHasher(), Default(), flagKey, subject, percent)
	return on
}
