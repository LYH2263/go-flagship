package errors_test

import (
	"errors"
	"fmt"
	"testing"

	ierr "github.com/LYH2263/go-flagship/internal/errors"
)

func TestWrapIs(t *testing.T) {
	inner := fmt.Errorf("bad attr")
	w := ierr.WrapErr(ierr.ErrInvalidRule, inner)
	if !errors.Is(w, ierr.ErrInvalidRule) {
		t.Fatalf("Is sentinel: %v", w)
	}
	if !errors.Is(w, inner) {
		t.Fatalf("Is inner: %v", w)
	}
}
