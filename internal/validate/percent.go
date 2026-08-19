package validate

import ierr "github.com/LYH2263/go-flagship/internal/errors"

func Percent(p int) error {
	if p < 0 || p > 100 {
		return ierr.ErrBadPercent
	}
	return nil
}
