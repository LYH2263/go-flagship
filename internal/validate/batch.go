package validate

import (
	"context"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
)

func Keys(ctx context.Context, keys []string) error {
	if ctx == nil {
		ctx = context.Background()
	}
	for _, k := range keys {
		if err := ctx.Err(); err != nil {
			return ierr.WrapErr(ierr.ErrCanceled, err)
		}
		if err := FlagKey(k); err != nil {
			return err
		}
	}
	return nil
}
