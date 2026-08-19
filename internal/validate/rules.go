package validate

import "github.com/LYH2263/go-flagship/internal/rule"

func Rules(specs []rule.Spec) error {
	return rule.ValidateAll(specs)
}
