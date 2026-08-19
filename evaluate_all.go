package flagship

import (
	"context"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
	"github.com/LYH2263/go-flagship/internal/validate"
)

func (s *Ship) EvaluateAll(ctx context.Context, keys []string, a EvalAttrs) (map[string]Decision, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	out := make(map[string]Decision, len(keys))
	for _, key := range keys {
		// 正确性钩子：每次迭代检查取消。
		if err := ctx.Err(); err != nil {
			return nil, ierr.WrapErr(ErrCanceled, err)
		}
		if err := validate.FlagKey(key); err != nil {
			return nil, err
		}
		dec, err := s.EvaluateContext(ctx, key, a)
		if err != nil {
			return nil, err
		}
		out[key] = dec
	}
	return out, nil
}
