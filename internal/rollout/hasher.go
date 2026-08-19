package rollout

import "hash/fnv"

type Hasher interface {
	Bucket(flagKey, subject string) int
}

type FNVHasher struct{}

func DefaultHasher() Hasher { return FNVHasher{} }

func (FNVHasher) Bucket(flagKey, subject string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(flagKey))
	_, _ = h.Write([]byte{0})
	_, _ = h.Write([]byte(subject))
	return int(h.Sum32() % 100)
}
