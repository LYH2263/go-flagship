package validate

import (
	ierr "github.com/LYH2263/go-flagship/internal/errors"
	"github.com/LYH2263/go-flagship/internal/flag"
)

func FlagKey(key string) error {
	if key == "" {
		return ierr.ErrEmptyKey
	}
	if !flag.IsValidKey(key) {
		return ierr.ErrInvalidKey
	}
	return nil
}
