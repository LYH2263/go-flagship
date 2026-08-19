package eval

import (
	"github.com/LYH2263/go-flagship/internal/rollout"
	"github.com/LYH2263/go-flagship/internal/rule"
)

func Run(in Input) Result {
	res := Result{Key: in.Key, Bucket: -1}
	if !in.Enabled {
		res.Reason = "disabled"
		return res
	}
	specs := make([]rule.Spec, 0, len(in.Rules))
	for _, r := range in.Rules {
		specs = append(specs, rule.Spec{Attr: r.Attr, Op: r.Op, Values: r.Values})
	}
	matched := rule.MatchAll(specs, in.Attrs)
	res.Matched = matched
	if !matched {
		res.Reason = "rules_miss"
		return res
	}
	subject := rollout.SubjectFromAttrs(in.Attrs)
	bucket, on := rollout.Decide(in.Hasher, in.Rollout, in.Key, subject, in.Percent)
	res.Bucket = bucket
	res.On = on
	if on {
		res.Reason = "rollout_on"
	} else {
		res.Reason = "rollout_off"
	}
	return res
}
