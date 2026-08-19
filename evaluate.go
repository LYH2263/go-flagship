package flagship

import (
	"context"

	"github.com/LYH2263/go-flagship/internal/attrs"
	ierr "github.com/LYH2263/go-flagship/internal/errors"
	"github.com/LYH2263/go-flagship/internal/eval"
	"github.com/LYH2263/go-flagship/internal/rollout"
)

func (s *Ship) Evaluate(key string, a EvalAttrs) (Decision, error) {
	return s.EvaluateContext(context.Background(), key, a)
}

func (s *Ship) EvaluateContext(ctx context.Context, key string, a EvalAttrs) (Decision, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	// 正确性钩子：入口检查 ctx.Err()。
	if err := ctx.Err(); err != nil {
		return Decision{}, ierr.WrapErr(ErrCanceled, err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// nil-guard：hasher/rollout 为空时安装默认。
	hasher := s.hasher
	if hasher == nil {
		hasher = rollout.DefaultHasher()
		s.hasher = hasher
	}
	ro := s.rollout
	if ro == nil {
		ro = rollout.Default()
		s.rollout = ro
	}

	rec, ok := s.table.Get(key)
	if !ok {
		return Decision{}, ErrNotFound
	}

	attrCopy := attrs.CloneMap(map[string]string(a))
	rules := fromFlagRules(rec.Rules)
	dec := eval.Run(eval.Input{
		Key:     key,
		Enabled: rec.Enabled,
		Percent: rec.Percent,
		Rules:   toEvalRules(rules),
		Attrs:   attrCopy,
		Hasher:  hasher,
		Rollout: ro,
	})
	if err := ctx.Err(); err != nil {
		return Decision{}, ierr.WrapErr(ErrCanceled, err)
	}
	s.evaluates++
	if dec.On {
		s.onCount++
	} else {
		s.offCount++
	}
	return Decision{
		Key:     dec.Key,
		On:      dec.On,
		Reason:  dec.Reason,
		Bucket:  dec.Bucket,
		Matched: dec.Matched,
	}, nil
}

func toEvalRules(in []Rule) []eval.Rule {
	out := make([]eval.Rule, 0, len(in))
	for _, r := range in {
		out = append(out, eval.Rule{
			Attr:   r.Attr,
			Op:     string(r.Op),
			Values: append([]string(nil), r.Values...),
		})
	}
	return out
}
