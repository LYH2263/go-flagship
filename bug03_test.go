package flagship_test

import (
	"errors"
	"testing"

	"github.com/LYH2263/go-flagship"
)

func TestBug03_EvaluateAfterCloseNoPanic(t *testing.T) {
	s := flagship.New()
	if err := s.Upsert(flagship.FlagDef{Key: "x", Enabled: true, Percent: 100}); err != nil {
		t.Fatal(err)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if rec := recover(); rec != nil {
			t.Fatalf("Evaluate panicked after Close: %v", rec)
		}
	}()
	_, err := s.Evaluate("x", flagship.EvalAttrs{"user": "u1"})
	if err == nil {
		t.Fatal("expected error after Close")
	}
	if !errors.Is(err, flagship.ErrClosed) && !errors.Is(err, flagship.ErrNilStore) {
		t.Fatalf("want ErrClosed/ErrNilStore, got %v", err)
	}
}
