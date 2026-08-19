package eval

import "github.com/LYH2263/go-flagship/internal/rollout"

type Rule struct {
	Attr   string
	Op     string
	Values []string
}

type Input struct {
	Key     string
	Enabled bool
	Percent int
	Rules   []Rule
	Attrs   map[string]string
	Hasher  rollout.Hasher
	Rollout rollout.Rollout
}

type Result struct {
	Key     string
	On      bool
	Reason  string
	Bucket  int
	Matched bool
}
