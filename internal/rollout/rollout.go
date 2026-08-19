package rollout

type Rollout interface {
	InRollout(bucket, percent int) bool
}

type PercentRollout struct{}

func Default() Rollout { return PercentRollout{} }

func (PercentRollout) InRollout(bucket, percent int) bool {
	if percent <= 0 {
		return false
	}
	if percent >= 100 {
		return true
	}
	if bucket < 0 {
		bucket = 0
	}
	if bucket > 99 {
		bucket = 99
	}
	return bucket < percent
}
