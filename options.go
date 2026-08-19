package flagship

import (
	"github.com/LYH2263/go-flagship/internal/clock"
	"github.com/LYH2263/go-flagship/internal/rollout"
)

type Option func(*Ship)

func WithClock(c clock.Clock) Option {
	return func(s *Ship) {
		if c != nil {
			s.clk = c
		}
	}
}

func WithPersistPath(path string) Option {
	return func(s *Ship) { s.persistPath = path }
}

func WithMaxFlags(n int) Option {
	return func(s *Ship) {
		if n > 0 {
			s.maxFlags = n
		}
	}
}

func WithHasher(h rollout.Hasher) Option {
	return func(s *Ship) { s.hasher = h }
}

func WithRollout(r rollout.Rollout) Option {
	return func(s *Ship) { s.rollout = r }
}

func WithDefaultPercent(p int) Option {
	return func(s *Ship) {
		if p >= 0 && p <= 100 {
			s.defaultPercent = p
		}
	}
}

func WithStrictRules(v bool) Option {
	return func(s *Ship) { s.strictRules = v }
}
