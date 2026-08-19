package rule

import (
	"fmt"
	"strings"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
)

func Validate(s Spec) error {
	if strings.TrimSpace(s.Attr) == "" {
		return ierr.Wrap(ierr.ErrInvalidRule, "empty attr")
	}
	op := strings.ToLower(s.Op)
	switch op {
	case "eq", "ne", "in", "notin", "exists", "missing":
	default:
		return ierr.Wrap(ierr.ErrInvalidRule, "unknown op "+s.Op)
	}
	if op == "eq" || op == "ne" || op == "in" || op == "notin" {
		if len(s.Values) == 0 {
			return ierr.Wrap(ierr.ErrInvalidRule, "values required for "+op)
		}
	}
	return nil
}

func ValidateAll(specs []Spec) error {
	for i, s := range specs {
		if err := Validate(s); err != nil {
			return ierr.WrapErr(ierr.ErrInvalidRule, fmt.Errorf("rule[%d]: %w", i, err))
		}
	}
	return nil
}

func AsInvalid(msg string) error {
	return ierr.Wrap(ierr.ErrInvalidRule, msg)
}
