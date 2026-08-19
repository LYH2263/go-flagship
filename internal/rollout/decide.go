package rollout

func Decide(h Hasher, r Rollout, flagKey, subject string, percent int) (bucket int, on bool) {
	if h == nil {
		h = DefaultHasher()
	}
	if r == nil {
		r = Default()
	}
	bucket = h.Bucket(flagKey, subject)
	on = r.InRollout(bucket, ClampPercent(percent))
	return bucket, on
}
