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
		if err := FlagKey(k); err != nil {
			return err
		}
	}
	return nil
}
